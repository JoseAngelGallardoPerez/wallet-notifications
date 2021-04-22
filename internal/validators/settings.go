package validators

import (
	"github.com/Confialink/wallet-notifications/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// UpdateSettingsValidator is
type UpdateSettingsValidator struct {
	Data []*db.Settings `json:"data" binding:"required,dive,required"`
}

// BindJSON binding from JSON
func (s *UpdateSettingsValidator) BindJSON(c *gin.Context) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	err := c.ShouldBindWith(s, b)
	if err != nil {
		return err
	}
	return nil
}
