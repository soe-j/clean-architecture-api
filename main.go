package main

import (
	"github.com/gin-gonic/gin"
	"github.com/soe-j/clean-architecture-api/src/drivers"
	"github.com/soe-j/clean-architecture-api/src/usecases"
)

func main() {
	r := gin.Default()

	h := drivers.GinHandler{}

	s := usecases.Service{}

	r.GET("/users", h.JSON(s.Index()))

	r.Run()
}
