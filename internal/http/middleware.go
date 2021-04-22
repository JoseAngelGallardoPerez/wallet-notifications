package http

import (
	"net/http"

	"github.com/Confialink/wallet-notifications/internal/service/auth"
	"github.com/Confialink/wallet-notifications/internal/service/users"
	"github.com/Confialink/wallet-pkg-acl"
	"github.com/Confialink/wallet-pkg-env_config"
	pbusers "github.com/Confialink/wallet-users/rpc/proto/users"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CorsMiddleware cors middleware
func CorsMiddleware(cfg *env_config.Cors) gin.HandlerFunc {

	corsConfig := cors.DefaultConfig()

	corsConfig.AllowMethods = cfg.Methods
	for _, origin := range cfg.Origins {
		if origin == "*" {
			corsConfig.AllowAllOrigins = true
		}
	}
	if !corsConfig.AllowAllOrigins {
		corsConfig.AllowOrigins = cfg.Origins
	}
	corsConfig.AllowHeaders = cfg.Headers

	return cors.New(corsConfig)
}

func AdminOnly(c *gin.Context) {
	user, exist := c.Get("_user")
	if !exist {
		c.Status(http.StatusForbidden)
		c.Abort()
		return
	}

	role := acl.RolesHelper.FromName((user.(*pbusers.User)).RoleName)
	if role < acl.Admin {
		c.Status(http.StatusForbidden)
		c.Abort()
		return
	}
}

type PermissionChecker struct {
	authService  *auth.Service
	usersService *users.Service
}

func NewMwPermissionChecker(authService *auth.Service, usersService *users.Service) *PermissionChecker {
	return &PermissionChecker{authService, usersService}
}

// check dynamic permission for resource
func (s *PermissionChecker) CanDynamic(action string, resourceName string, resource interface{}) func(*gin.Context) {
	return func(c *gin.Context) {
		user, exist := c.Get("_user")
		if !exist {
			c.Status(http.StatusForbidden)
			c.Abort()
			return
		}

		if !s.authService.CanDynamic(user.(*pbusers.User), action, resourceName, resource) {
			c.Status(http.StatusForbidden)
			c.Abort()
			return
		}
	}
}

// check dynamic permission for resource UserSettings
func (s *PermissionChecker) CanDynamicUserSettings(action string, resourceName string) func(*gin.Context) {
	return func(c *gin.Context) {
		user, exist := c.Get("_user")
		if !exist {
			c.Status(http.StatusForbidden)
			c.Abort()
			return
		}
		uid := c.Params.ByName("uid")
		if uid == "" {
			c.Status(http.StatusForbidden)
			c.Abort()
			return
		}

		requestedUser, err := s.usersService.GetByUID(uid)
		if err != nil {
			c.Status(http.StatusForbidden)
			c.Abort()
			return
		}

		if !s.authService.CanDynamic(user.(*pbusers.User), action, resourceName, requestedUser) {
			c.Status(http.StatusForbidden)
			c.Abort()
			return
		}
	}
}
