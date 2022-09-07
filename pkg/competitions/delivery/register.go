package delivery

import (
	"github.com/ArdentK/db-cp-final/pkg/competitions"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, usecase competitions.CompUseCase) {
	h := newHandler(usecase)

	router.GET("/competitions", h.Index)
	router.GET("/analytics", h.Analytics)
}
