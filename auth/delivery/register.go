package delivery

import (
	"github.com/ArdentK/db-cp-final/auth"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, usecase auth.UseCase) {
	h := NewHandler(usecase)

	router.POST("/sign-up", h.SignUp)
	router.POST("/sign-in", h.SignIn)
}
