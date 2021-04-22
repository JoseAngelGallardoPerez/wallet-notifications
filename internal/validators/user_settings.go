package validators

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// UpdateUserSettingsValidator is
type UpdateUserSettingsValidator struct {
	Data map[string]string `json:"data" binding:"required,dive,required"`
}

// BindJSON binding from JSON
func (s *UpdateUserSettingsValidator) BindJSON(c *gin.Context) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	err := c.ShouldBindWith(s, b)
	if err != nil {
		return err
	}
	return nil
}
