package http

import (
	"net/http"

	env_config "github.com/Confialink/wallet-pkg-env_config"
	"github.com/Confialink/wallet-pkg-errors"
	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"

	"github.com/Confialink/wallet-notifications/internal/authentication"
	"github.com/Confialink/wallet-notifications/internal/service/auth"
	"github.com/Confialink/wallet-notifications/internal/service/users"
	"github.com/Confialink/wallet-notifications/internal/version"
)

var r *gin.Engine

type Router struct {
	service          *Service
	corsConfig       *env_config.Cors
	logger           log15.Logger
	pushTokenHandler *PushTokenHandler
}

// NewRouter creates new router
func NewRouter(service *Service, corsConfig *env_config.Cors, logger log15.Logger, pushTokenHandler *PushTokenHandler) *Router {
	return &Router{service, corsConfig, logger, pushTokenHandler}
}

// RegisterRoutes is where you can register all of the routes for an service.
func (router Router) RegisterRoutes() *gin.Engine {

	// Creates a gin router with default middleware:
	r = gin.Default()

	r.GET("/notifications/health-check", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	r.GET("/notifications/build", func(c *gin.Context) {
		c.JSON(http.StatusOK, version.BuildInfo)
	})

	errorsMiddleware := errors.ErrorHandler(router.logger.New("Middleware", "Errors"))

	r.Use(
		gin.Recovery(),
		gin.Logger(),
		CorsMiddleware(router.corsConfig),
		errorsMiddleware,
	)

	apiGroup := r.Group("notifications")

	/*
	 |---------------------------------------------------
	 | Private router group
	 |---------------------------------------------------
	*/
	privateGroup := apiGroup.Group("/private", authentication.Middleware(router.logger))
	{
		v1Group := privateGroup.Group("/v1")
		{
			permChecker := NewMwPermissionChecker(auth.NewService(auth.NewPermissionsChecker()), users.NewService())
			mwViewSettingsPerm := permChecker.CanDynamic(auth.ActionUpdate, auth.Resources, auth.ViewSettings)
			mwModifySettingsPerm := permChecker.CanDynamic(auth.ActionUpdate, auth.Resources, auth.ModifySettings)

			settingsGroup := v1Group.Group("/settings", AdminOnly)
			{
				// GET /notifications/private/v1/settings
				settingsGroup.GET("", mwViewSettingsPerm, router.service.GetSettingsHandler)
				// PUT /notifications/private/v1/settings
				settingsGroup.PUT("", mwModifySettingsPerm, router.service.UpdateSettingsHandler)
				// GET /notifications/private/v1/settings/tokens
				settingsGroup.GET("/tokens", mwViewSettingsPerm, router.service.GetTokens)
			}

			userSettingsGroup := v1Group.Group("/user-settings")
			{
				// GET /notifications/private/v1/user-settings
				userSettingsGroup.GET("", router.service.GetUserSettingsHandler)
				// PUT /notifications/private/v1/user-settings
				userSettingsGroup.PUT("", router.service.UpdateUserSettingsHandler)

				mwReadPermissions := permChecker.CanDynamicUserSettings(auth.ActionRead, auth.ResourceUserSettings)
				mwUpdatePermissions := permChecker.CanDynamicUserSettings(auth.ActionUpdate, auth.ResourceUserSettings)
				// GET /notifications/private/v1/user-settings
				userSettingsGroup.GET("/:uid", AdminOnly, mwReadPermissions, router.service.GetUserSettingsHandler)
				// PUT /notifications/private/v1/user-settings
				userSettingsGroup.PUT("/:uid", AdminOnly, mwUpdatePermissions, router.service.UpdateUserSettingsHandler)
			}

			templatesGroup := v1Group.Group("/templates", AdminOnly)
			{
				// GET /notifications/private/v1/templates/:scope
				templatesGroup.GET("/:scope", mwViewSettingsPerm, router.service.GetTemplatesHandler)
				// PUT /notifications/private/v1/templates
				templatesGroup.PUT("", mwModifySettingsPerm, router.service.UpdateTemplatesHandler)
			}

			notifiersGroup := v1Group.Group("/notifiers", AdminOnly, mwViewSettingsPerm)
			{
				// GET /notifications/private/v1/notifiers/plivo/details
				notifiersGroup.GET("/:provider/details", router.service.GetNotifierProviderDetails)
			}

			testGroup := v1Group.Group("/test", AdminOnly, mwViewSettingsPerm)
			{
				// POST /notifications/private/v1/test/smtp
				testGroup.POST("smtp", router.service.TestSMTPSettingsHandler)
				// POST /notifications/private/v1/test/events
				testGroup.POST("events", router.service.TestEventHandler)
				// POST /notifications/private/v1/test/sms
				testGroup.POST("sms", router.service.TestSMSSettingsHandler)
				// POST /notifications/private/v1/test/push
				testGroup.POST("push", router.pushTokenHandler.Test)
			}

			pushTokenGroup := v1Group.Group("/push-tokens")
			{
				// POST /notifications/private/v1/push-tokens
				pushTokenGroup.POST("", router.pushTokenHandler.CreateOrUpdate)
				// POST /notifications/private/v1/push-tokens/delete
				// We use POST instead of DELETE because the body is not empty.
				pushTokenGroup.POST("delete", router.pushTokenHandler.Delete)
				// GET /notifications/private/v1/push-tokens/user/{uid}
				pushTokenGroup.GET("/user/:uid", AdminOnly, router.pushTokenHandler.List)
			}
		}
	}

	/*
	 |---------------------------------------------------
	 | Public router group
	 |---------------------------------------------------
	*/
	publicGroup := apiGroup.Group("/public")
	{
		v1PublicGroup := publicGroup.Group("/v1")
		{
			settingsPublicGroup := v1PublicGroup.Group("/settings")
			{
				// GET /notifications/public/v1/settings/email-from
				settingsPublicGroup.GET("/email-from", router.service.GetEmailFromHandler)
			}
		}
	}

	// If route not found returns StatusNotFound
	r.NoRoute(router.service.NotFoundHandler)

	// Handle OPTIONS request
	r.OPTIONS("/*cors", router.service.OptionsHandler)

	return r
}
