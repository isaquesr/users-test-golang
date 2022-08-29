package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaquesr/users-test-golang/auth"
	"github.com/isaquesr/users-test-golang/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct {
	useCase auth.UseCase
}

func NewHandler(useCase auth.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type SignInput struct {
	ID       primitive.ObjectID `json:"id"`
	Name     string             `json:"name"`
	Password string             `json:"password"`
	Address  string             `json:"address"`
	Age      int32              `json:"age"`
	Email    string             `json:"email"`
}

type SignIn struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (h *Handler) SignUp(c *gin.Context) {
	user := &domain.User{
		ID:       primitive.NewObjectID(),
		Name:     "name",
		Password: "password",
		Email:    "email",
		Age:      20,
		Address:  "address",
	}
	if err := c.BindJSON(user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.useCase.SignUp(c.Request.Context(), user); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

type signInResponse struct {
	Token string `json:"token"`
}

func (h *Handler) SignIn(c *gin.Context) {
	in := new(SignIn)

	if err := c.BindJSON(in); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := h.useCase.SignIn(c.Request.Context(), in.Name, in.Password)
	if err != nil {
		if err == auth.ErrUserNotFound {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, signInResponse{Token: token})
}
