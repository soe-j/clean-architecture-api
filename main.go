package main

import (
	"github.com/gin-gonic/gin"
	"github.com/soe-j/clean-architecture-api/src/entities"
)

func main() {
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

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"user":    user,
			"company": company,
		})
	})
	r.Run()
}
