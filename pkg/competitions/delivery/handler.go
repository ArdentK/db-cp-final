package delivery

import (
	"net/http"

	"github.com/ArdentK/db-cp-final/pkg/competitions"

	"github.com/gin-gonic/gin"
)

type request struct {
	ID int `json:"id"`
}

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
	useCase competitions.CompUseCase
}

func newHandler(useCase competitions.CompUseCase) *handler {
	return &handler{
		useCase: useCase,
	}
}

// TODO
// Добавить распарсивание параметров!!!
func (h *handler) Index(c *gin.Context) {
	comp, err := h.useCase.List(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newResponse(STATUS_ERROR, err.Error()))
		return
	}

	c.JSON(http.StatusOK, comp)
}

func (h *handler) Analytics(c *gin.Context) {
	c.JSON(http.StatusOK, "ANALYTICS")
}

func (h *handler) Get(c *gin.Context) {
	inp := new(request)
	err := c.BindJSON(inp)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newResponse(STATUS_ERROR, err.Error()))
		return
	}

	comp, err := h.useCase.FindByID(c.Request.Context(), inp.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, newResponse(STATUS_ERROR, err.Error()))
		return
	}

	c.JSON(http.StatusOK, comp)
}
