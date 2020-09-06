package main

import (
	"github.com/gin-gonic/gin"
	"github.com/soe-j/clean-architecture-api/src/drivers"
)

func main() {
	r := gin.Default()

	h := drivers.GinHandler{}

	r.GET("/users", h.JSON())

	r.Run()
}
