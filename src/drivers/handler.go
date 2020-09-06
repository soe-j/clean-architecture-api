package drivers

import (
	"github.com/gin-gonic/gin"
	"github.com/soe-j/clean-architecture-api/src/usecases"
)

// GinHandler ...
type GinHandler struct{}

// JSON ...
func (h *GinHandler) JSON(data usecases.ServiceResponse) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H(data))
	}
}
