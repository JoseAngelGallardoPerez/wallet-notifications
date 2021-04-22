package http

import (
	"net/http"

	errors "github.com/Confialink/wallet-pkg-errors"
	list_params "github.com/Confialink/wallet-pkg-list_params"
	userpb "github.com/Confialink/wallet-users/rpc/proto/users"
	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"

	"github.com/Confialink/wallet-notifications/internal/db"
	"github.com/Confialink/wallet-notifications/internal/db/models"
	"github.com/Confialink/wallet-notifications/internal/errcodes"
	"github.com/Confialink/wallet-notifications/internal/service/providers"
	"github.com/Confialink/wallet-notifications/internal/service/providers/firebase"
	push_token "github.com/Confialink/wallet-notifications/internal/service/push-token"
	"github.com/Confialink/wallet-notifications/internal/validators"
)

type PushTokenHandler struct {
	resp    *ResponseService
	service *push_token.Service
	repo    db.RepositoryInterface
	logger  log15.Logger
}

func NewPushTokenHandler(resp *ResponseService, service *push_token.Service, repo db.RepositoryInterface, logger log15.Logger) *PushTokenHandler {
	return &PushTokenHandler{resp, service, repo, logger.New("handler", "PushTokenHandler")}
}

func (s *PushTokenHandler) CreateOrUpdate(c *gin.Context) {
	user, ok := c.Get("_user")
	if !ok {
		s.resp.ErrorResponse(c, Forbidden, "User not found")
		return
	}

	form := &validators.AddPushToken{}
	if err := c.ShouldBind(form); err != nil {
		s.resp.ValidatorErrorResponse(c, UnprocessableEntity, err)
		return
	}

	res, err := s.service.AddOrUpdate(form, user.(*userpb.User).UID)
	if err != nil {
		s.logger.New("method", "CreateOrUpdate").Error("cannot create or update push token", "err", err)
		s.resp.ErrorResponse(c, UnprocessableEntity, "Bad Request")
		return
	}

	s.resp.OkResponse(c, res)
}

func (s *PushTokenHandler) Delete(c *gin.Context) {
	form := &validators.DeletePushToken{}
	if err := c.ShouldBind(form); err != nil {
		s.resp.ValidatorErrorResponse(c, UnprocessableEntity, err)
		return
	}

	if err := s.service.RemovePushToken(form.PushToken); err != nil {
		s.logger.New("method", "Delete").Error("cannot delete push token", "err", err)
		s.resp.ErrorResponse(c, UnprocessableEntity, "Bad Request")
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (s *PushTokenHandler) Test(c *gin.Context) {
	logger := s.logger.New("method", "Test")
	var form struct {
		PushToken string `json:"pushToken" binding:"required,max=255"`
		Title     string `json:"title" binding:"required,max=255"`
		Body      string `json:"body" binding:"required,max=255"`
	}
	if err := c.ShouldBindJSON(&form); err != nil {
		s.resp.ValidatorErrorResponse(c, UnprocessableEntity, err)
		return
	}
	pushToken, err := s.service.FindOne(form.PushToken)
	if err != nil {
		logger.Error("cannot find push token", "err", err)
		s.resp.ErrorResponse(c, UnprocessableEntity, "")
		return
	}
	if !pushToken.IsExists() {
		logger.Error("push token not found", "err", err)
		s.resp.ErrorResponse(c, UnprocessableEntity, "Push token not found")
		return
	}

	settings, err := s.repo.GetAll()
	if err != nil {
		logger.Error("cannot receive settings", "err", err)
		s.resp.ErrorResponse(c, UnprocessableEntity, "Bad Request")
		return
	}

	for _, setting := range settings {
		// We enable push status
		if setting.Name == db.KeyPushStatus {
			setting.Value = "enabled"
		}
	}

	client, err := firebase.Load(settings, s.logger)
	if err != nil {
		logger.Error("cannot load firebase", "err", err)
		s.resp.ErrorResponse(c, UnprocessableEntity, "Bad Request")
		return
	}

	firebaseClient, ok := client.(providers.ByPush)
	if !ok {
		logger.Error("invalid firebase client", "err", err)
		s.resp.ErrorResponse(c, UnprocessableEntity, "Bad Request")
		return
	}

	if err := firebaseClient.SetBody(form.Body).SetSubject(form.Title).To([]string{form.PushToken}).Send(); err != nil {
		typedErr := &errors.PublicError{
			Code:       errcodes.ErrorSendingTestPush,
			Title:      err.Error(),
			HttpStatus: http.StatusBadRequest,
		}
		errors.AddErrors(c, typedErr)
		return
	}

	// Returns a "204 StatusNoContent" response
	c.JSON(http.StatusNoContent, nil)
}

func (s *PushTokenHandler) List(c *gin.Context) {
	uid := c.Params.ByName("uid")

	listParams := s.listParams(c.Request.URL.RawQuery, uid)
	if ok, errorsList := listParams.Validate(); !ok {
		errcodes.AddErrorMeta(c, CodeInvalidQueryParameters, errorsList)
		return
	}
	listParams.AddFilter("uid", []string{uid})
	res, err := s.service.List(listParams)
	if err != nil {
		s.logger.New("method", "List").Error("cannot get list of push tokens", "err", err)
		s.resp.ErrorResponse(c, UnprocessableEntity, "Bad Request")
		return
	}

	count, err := s.service.Count(listParams)
	if err != nil {
		s.logger.New("method", "List").Error("cannot count push tokens", "err", err)
		s.resp.ErrorResponse(c, UnprocessableEntity, "Bad Request")
		return
	}

	s.resp.OkResponseWithPagination(c, res, listParams, count)
}

func (s *PushTokenHandler) listParams(query, uid string) *list_params.ListParams {
	params := list_params.NewListParamsFromQuery(query, models.PushToken{})
	params.AllowPagination()
	params.AllowFilters([]string{"os", list_params.FilterLike("name"), "deviceId"})
	params.AllowSortings([]string{
		"os", "name", "deviceId", "createdAt", "updatedAt",
	})

	return params
}
