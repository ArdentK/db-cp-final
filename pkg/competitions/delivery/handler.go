package delivery

import (
	"net/http"

	"github.com/ArdentK/db-cp-final/pkg/competitions"

	"github.com/gin-gonic/gin"
)

type handler struct {
	useCase competitions.CompUseCase
}

func newHandler(useCase competitions.CompUseCase) *handler {
	return &handler{
		useCase: useCase,
	}
}

func (h *handler) List(c *gin.Context) {
	c.JSON(http.StatusOK, "OLOLO")
}
