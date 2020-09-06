package drivers

import (
	"github.com/gin-gonic/gin"
	"github.com/soe-j/clean-architecture-api/src/entities"
)

// GinHandler ...
type GinHandler struct{}

// JSON ...
func (h *GinHandler) JSON() gin.HandlerFunc {

	company := &entities.Company{
		ID:      1,
		Code:    "GITHUB",
		Name:    "GitHub",
		Address: "USA",
	}
	user := &entities.User{
		ID:        1,
		Email:     "soe-j@github.com",
		Name:      "soe-j",
		CompanyID: company.ID,
	}

	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"user":    user,
			"company": company,
		})
	}
}
