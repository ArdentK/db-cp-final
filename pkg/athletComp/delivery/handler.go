package delivery

import (
	"net/http"

	athletcomp "github.com/ArdentK/db-cp-final/pkg/athletComp"
	"github.com/gin-gonic/gin"
)

const (
	STATUS_OK    = "ok"
	STATUS_ERROR = "error"
)

type response struct {
	Status string `json:"status"`
	Msg    string `json:"message,omitempty"`
}

func newResponse(status, msg string) *response {
	return &response{
		Status: status,
		Msg:    msg,
	}
}

type handler struct {
	useCase athletcomp.ACUseCase
}

func newHandler(useCase athletcomp.ACUseCase) *handler {
	return &handler{
		useCase: useCase,
	}
}

//TODO
func (h *handler) NewRow(c *gin.Context) {
	c.JSON(http.StatusOK, "ADD ROW")
}
