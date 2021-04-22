package http

import (
	"errors"
	"net/http"

	errPkg "github.com/Confialink/wallet-pkg-errors"
	"github.com/Confialink/wallet-users/rpc/proto/users"
	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"

	"github.com/Confialink/wallet-notifications/internal/db"
	"github.com/Confialink/wallet-notifications/internal/errcodes"
	"github.com/Confialink/wallet-notifications/internal/eventmanager"
	"github.com/Confialink/wallet-notifications/internal/service"
	"github.com/Confialink/wallet-notifications/internal/service/providers"
	"github.com/Confialink/wallet-notifications/internal/service/providers/plivo"
	"github.com/Confialink/wallet-notifications/internal/validators"
)

// Service is base struct
type Service struct {
	repo       db.RepositoryInterface
	resp       *ResponseService
	notify     *service.Notifier
	em         *eventmanager.EventManager
	serializer *Serializer
	logger     log15.Logger
}

// NewService creates new service
func NewService(
	repo db.RepositoryInterface,
	resp *ResponseService,
	notify *service.Notifier,
	em *eventmanager.EventManager,
	serializer *Serializer,
	logger log15.Logger,
) *Service {
	return &Service{repo, resp, notify, em, serializer, logger}
}

// OptionsHandler handle options request
func (s *Service) OptionsHandler(c *gin.Context) {}

// NotFoundHandler returns 404 NotFound
func (s *Service) NotFoundHandler(c *gin.Context) {
	// Returns a "404 StatusNotFound" response
	s.resp.NotFoundResponse(c)
	return
}

// GetCurrentUserId return current user id
func (s *Service) GetCurrentUserId(c *gin.Context) string {
	uid := c.Params.ByName("uid")
	if uid == "" {
		user, _ := c.Get("_user")
		curentUser := (user.(*users.User))
		uid = curentUser.UID
	}
	return uid
}

// GetUserSettingsHandler returns list of settings
func (s *Service) GetUserSettingsHandler(c *gin.Context) {
	uid := s.GetCurrentUserId(c)

	settingsList := service.GetNotificationNamesFromUserSettings()

	var settings map[string]string = map[string]string{}

	for _, item := range settingsList {

		setting, err := s.repo.FindUserOptionByUIDAndName(uid, item)
		if nil != err {
			// Returns a "400 StatusBadRequest" response
			s.resp.ErrorResponse(c, CanNotRetrieveCollection, "Bad Request")
			return
		}
		settings[item] = setting.IsActive
	}

	// Returns a "200 OK" response
	s.resp.OkResponse(c, settings)
}

// UpdateSettingsHandler updates settings
func (s *Service) UpdateUserSettingsHandler(c *gin.Context) {
	uid := s.GetCurrentUserId(c)

	// Checks if the query entry is valid
	validator := validators.UpdateUserSettingsValidator{}
	if err := validator.BindJSON(c); err != nil {
		// Returns a "422 StatusUnprocessableEntity" response
		s.resp.ValidatorErrorResponse(c, CanNotUpdateCollection, err)
		return
	}

	settingsList := service.GetNotificationNamesFromUserSettings()
	for _, item := range settingsList {
		err := s.repo.UpdateUserSetting(&db.UserSettings{
			NotificationName: item,
			UID:              uid,
			IsActive:         validator.Data[item],
		})

		if err != nil {
			// Returns a "500 StatusInternalServerError" response
			s.resp.ValidatorErrorResponse(c, CanNotUpdateCollection, err)
			return
		}
	}

	// Returns a "204 StatusNoContent" response
	c.JSON(http.StatusNoContent, nil)
}

// GetTokens returns tokens
func (s *Service) GetTokens(c *gin.Context) {
	html := "[UserName]<br/>" +
		"[Email]<br/>" +
		"[FirstName]<br/>" +
		"[LastName]<br/>" +
		"[Password]<br/>" +
		"[SiteName]<br/>" +
		"[SiteLoginURL]<br/>" +
		"[SiteURL]<br/>" +
		"[PasswordRecoveryURL]<br/>" +
		"[Logo] (Restricted to Mail signature)<br/>" +
		"[OneTimeLoginURL] (Restricted to New profile Welcome)<br/>" +
		"[SetPasswordOneTimeURL] (Restricted to New profile Welcome)<br/>" +
		"[PrivateMessageRecipient] (Restricted to Incoming Message)<br/>" +
		"[PrivateMessageAuthor] (Restricted to Incoming Message)<br/>" +
		"[PrivateMessageURL] (Restricted to Incoming Message)<br/>" +
		"[PrivateMessageRecipientEditURL] (Restricted to Incoming Message)<br/>" +
		"[Reason] (Restricted to Canceled Registration)<br/>" +
		"[Link] (Restricted to New News Created)<br/>" +
		"[DocumentName] (Restricted to New User Document)<br/>" +
		// "[Tan]<br/>" +
		"[VerificationLink]<br/>" +
		"[AccountNumber]"

	// Returns a "200 OK" response
	s.resp.OkResponse(c, html)
}

// GetSettingsHandler returns list of settings
func (s *Service) GetSettingsHandler(c *gin.Context) {
	settings, err := s.repo.GetAll()
	if nil != err {
		// Returns a "400 StatusBadRequest" response
		s.resp.ErrorResponse(c, CanNotRetrieveCollection, "Bad Request")
		return
	}
	// Returns a "200 OK" response
	s.resp.OkResponse(c, settings)
}

// GetSettingsHandler returns list of settings
func (s *Service) GetEmailFromHandler(c *gin.Context) {
	setting, err := s.repo.GetByName(db.KeyEmailFrom)
	if nil != err {
		s.resp.ErrorResponse(c, NotFound, "Not Found")
		return
	}

	s.resp.OkResponse(c, s.serializer.Serialize(setting))
}

// UpdateSettingsHandler updates settings
func (s *Service) UpdateSettingsHandler(c *gin.Context) {
	// Checks if the query entry is valid
	validator := validators.UpdateSettingsValidator{}
	if err := validator.BindJSON(c); err != nil {
		// Returns a "422 StatusUnprocessableEntity" response
		s.resp.ValidatorErrorResponse(c, UnprocessableEntity, err)
		return
	}
	// Update settings
	for _, item := range validator.Data {
		err := s.repo.FirstOrCreate(item)
		if err != nil {
			// Returns a "500 StatusInternalServerError" response
			s.resp.ErrorResponse(c, CanNotUpdateCollection, "Could not update settings")
			return
		}
	}
	// Returns a "204 StatusNoContent" response
	c.JSON(http.StatusNoContent, nil)
}

// GetTemplatesHandler returns list of templates
func (s *Service) GetTemplatesHandler(c *gin.Context) {
	scope := c.Params.ByName("scope")

	templates, err := s.repo.FindByScopeAndEditable(scope)
	if nil != err {
		// Returns a "400 StatusBadRequest" response
		s.resp.ErrorResponse(c, CanNotRetrieveCollection, "Bad Request")
		return
	}
	// Returns a "200 OK" response
	s.resp.OkResponse(c, templates)
}

// UpdateTemplatesHandler mass updates templates
func (s *Service) UpdateTemplatesHandler(c *gin.Context) {
	// Checks if the query entry is valid
	validator := validators.UpdateTemplatesValidator{}
	if err := validator.BindJSON(c); err != nil {
		// Returns a "422 StatusUnprocessableEntity" response
		s.resp.ValidatorErrorResponse(c, UnprocessableEntity, err)
		return
	}
	// Update templates
	for _, item := range validator.Data {
		err := s.repo.UpdateTemplate(item)
		if err != nil {
			// Returns a "500 StatusInternalServerError" response
			s.resp.ErrorResponse(c, CanNotUpdateCollection, "Could not update templates")
			return
		}
	}
	// Returns a "204 StatusNoContent" response
	c.JSON(http.StatusNoContent, nil)
}

// TestSMTPSettingsHandler sends test email
func (s *Service) TestSMTPSettingsHandler(c *gin.Context) {
	// Checks if the query entry is valid
	validator := validators.UpdateSettingsValidator{}
	if err := validator.BindJSON(c); err != nil {
		// Returns a "422 StatusUnprocessableEntity" response
		s.resp.ValidatorErrorResponse(c, UnprocessableEntity, err)
		return
	}

	settings := validator.Data

	testEmail, err := db.GetSettingValue("test_email", settings)
	if err != nil {
		// Returns a "400 StatusBadRequest" response
		s.resp.ErrorResponse(c, BadTestSMTPParams, err.Error())
		return
	}

	td := db.TemplateData{Email: testEmail}
	data := *service.NewCallData(
		"",
		[]service.Method{service.MethodEmail},
		&td,
	)

	data.To = testEmail

	template, err := s.repo.FindOneByTitleAndScope("Test email", db.ScopeAdmin)
	if err != nil {
		s.resp.ErrorResponse(c, BadTestSMTPParams, err.Error())
		return
	}

	template.ApplyTemplateData(data.TemplateData)

	receiver := service.NewReceiver(data.GetToByMethod(service.MethodEmail), service.MethodEmail)
	err = s.notify.SendToReceiver(template, receiver, data.TemplateData)
	if err != nil {
		s.resp.ErrorResponse(c, BadTestSMTPParams, "Test email failed - Please check SMTP settings.")
	}
	// Returns a "204 StatusNoContent" response
	c.JSON(http.StatusNoContent, nil)
}

// TestEventHandler is test handler
func (s *Service) TestEventHandler(c *gin.Context) {
	// Checks if the query entry is valid
	validator := validators.TestEventValidator{}
	if err := validator.BindJSON(c); err != nil {
		// Returns a "422 StatusUnprocessableEntity" response
		s.resp.ValidatorErrorResponse(c, UnprocessableEntity, err)
		return
	}

	event := validator.Data.EventName
	uid := validator.Data.To
	td := &db.TemplateData{
		OneTimeLoginURL:                "OneTimeLoginURLHere",
		PrivateMessageRecipient:        "PrivateMessageRecipientHere",
		PrivateMessageAuthor:           "PrivateMessageAuthorHere",
		PrivateMessageURL:              "PrivateMessageURLHere",
		PrivateMessageRecipientEditURL: "PrivateMessageRecipientEditURLHere",
		Reason:                         "ReasonHere",
		Link:                           "LinkHere",
		DocumentName:                   "DocumentNameHere",
		Tan:                            "TanHere",
		Password:                       "PasswordHere",
		VerificationLink:               "VerificationLinkHere",
	}

	notifiers := []service.Method{}

	data := service.NewCallData(uid, notifiers, td)

	s.em.Dispatch(event, *data)

	// Returns a "204 StatusNoContent" response
	c.JSON(http.StatusNoContent, nil)
}

// TestSMSSettingsHandler sends test SMS
func (s *Service) TestSMSSettingsHandler(c *gin.Context) {
	var f struct {
		Number   string `json:"number" binding:"required,max=20"`
		Provider string `json:"provider" binding:"omitempty,oneof=plivo twilio"`
	}
	if err := c.ShouldBindJSON(&f); err != nil {
		_ = c.Error(err)
		return
	}

	settings, err := s.repo.GetAll()
	if err != nil {
		_ = c.Error(err)
		return
	}

	if len(f.Provider) != 0 {
		for _, setting := range settings {
			if setting.Name == service.SettingNameSmsProvider {
				setting.Value = f.Provider
			}
		}
	}

	smsLoader, err := service.LoadSmsProvider(settings)
	if err != nil {
		_ = c.Error(err)
		return
	}

	smsProvider, err := smsLoader(settings, s.logger)
	if err != nil {
		_ = c.Error(err)
		return
	}

	typedProvider, ok := smsProvider.(providers.BySMS)
	if !ok {
		_ = c.Error(errors.New("invalid sms provider"))
		return
	}

	err = typedProvider.To(f.Number).SetBody("Test SMS").Send()
	if err != nil {
		if typedErr, ok := err.(errPkg.TypedError); ok {
			errPkg.AddErrors(c, typedErr)
			return
		}
		typedErr := &errPkg.PublicError{
			Code:       errcodes.ErrorSendingTestSms,
			Title:      err.Error(),
			HttpStatus: http.StatusBadRequest,
		}
		errPkg.AddErrors(c, typedErr)
		return
	}

	// Returns a "204 StatusNoContent" response
	c.JSON(http.StatusNoContent, nil)
}

// GetNotifierProviderDetails returns details of notifier provider
func (s *Service) GetNotifierProviderDetails(c *gin.Context) {
	provider := c.Params.ByName("provider")
	d, err := s.getProviderDetailsByName(provider)
	if err != nil {
		c.Error(err)
		return
	}

	// Returns a "200 StatusOK" response
	s.resp.OkResponse(c, d)
}

func (s *Service) getProviderDetailsByName(name string) (interface{}, error) {
	switch name {
	case "plivo":
		return s.getPlivoDetails()
	case "firebase":
		return nil, errors.New("can't get details for firebase")
	case "slack":
		return nil, errors.New("can't get details for slack")
	default:
		return nil, errors.New("unknown provider")
	}
}

func (s *Service) getPlivoDetails() (interface{}, error) {
	settings, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	client, err := plivo.Load(settings, s.logger)
	if err != nil {
		return nil, err
	}

	d, err := client.(*plivo.Client).GetAccountDetails()
	if err != nil {
		return nil, err
	}

	return struct {
		AccountType  string `json:"accountType,omitempty"`
		Address      string `json:"address,omitempty"`
		ApiId        string `json:"apiId,omitempty"`
		AuthId       string `json:"authId,omitempty"`
		AutoRecharge bool   `json:"autoRecharge,omitempty"`
		BillingMode  string `json:"billingMode,omitempty"`
		CashCredits  string `json:"cashCredits,omitempty"`
		City         string `json:"city,omitempty"`
		Name         string `json:"name,omitempty"`
		ResourceUri  string `json:"resourceUri,omitempty"`
		State        string `json:"state,omitempty"`
		Timezone     string `json:"timezone,omitempty"`
	}{
		d.AccountType,
		d.Address,
		d.ApiId,
		d.AuthId,
		d.AutoRecharge,
		d.BillingMode,
		d.CashCredits,
		d.City,
		d.Name,
		d.ResourceUri,
		d.State,
		d.Timezone,
	}, nil
}
