package delivery

import (
	"fmt"
	"net/http"

	"github.com/ArdentK/db-cp-final/auth"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase auth.UseCase
}

func NewHandler(useCase auth.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type signInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) SignUp(c *gin.Context) {
	inp := new(signInput)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.useCase.SignUp(c.Request.Context(), inp.Email, inp.Password); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	c.Status(http.StatusOK)
}

type signInResponse struct {
	Token string `json:"token"`
}

func (h *Handler) SignIn(c *gin.Context) {
	inp := new(signInput)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := h.useCase.SignIn(c.Request.Context(), inp.Email, inp.Password)
	if err != nil {
		if err == auth.ErrUserDoesNotExist {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, signInResponse{Token: token})
}
