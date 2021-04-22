package validators

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// TestEventValidator is
type TestEventValidator struct {
	Data struct {
		To        string `json:"to" binding:"max=255"`
		EventName string `json:"event_name" binding:"max=255"`
	} `json:"data"`
}

// BindJSON binding from JSON
func (s *TestEventValidator) BindJSON(c *gin.Context) error {
	b := binding.Default(c.Request.Method, c.ContentType())

	err := c.ShouldBindWith(s, b)
	if err != nil {
		return err
	}

	return nil
}
