package delivery

import (
	"net/http"

	"github.com/ArdentK/db-cp-final/auth"
	"github.com/ArdentK/db-cp-final/models"
	athletcomp "github.com/ArdentK/db-cp-final/pkg/athletComp"
	"github.com/gin-gonic/gin"
)

const (
	STATUS_OK    = "ok"
	STATUS_ERROR = "error"
)

type request struct {
	// IDAthlet int `json:"idathlet"`
	IDComp int `json:"idcomp"`
}

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
	// req := new(request)

	// err := c.BindJSON(req)
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, newResponse(STATUS_ERROR, err.Error()))
	// 	return
	// }

	// user := c.MustGet(auth.CtxUserKey).(*models.User)

	// err = h.useCase.Create(c.Request.Context(), user, req.IDComp)
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, newResponse(STATUS_ERROR, err.Error()))
	// 	return
	// }

	c.JSON(http.StatusOK, "ADD ROW")
}

func (h *handler) DelRow(c *gin.Context) {
	req := new(request)

	err := c.BindJSON(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newResponse(STATUS_ERROR, err.Error()))
		return
	}

	user := c.MustGet(auth.CtxUserKey).(*models.User)

	err = h.useCase.Delete(c.Request.Context(), user, req.IDComp)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, newResponse(STATUS_ERROR, err.Error()))
		return
	}

	c.JSON(http.StatusOK, "DEL ROW")
}

func (h *handler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, "TEST")
}
