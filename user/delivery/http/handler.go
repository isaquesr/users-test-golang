package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaquesr/users-test-golang/auth"
	"github.com/isaquesr/users-test-golang/domain"
	"github.com/isaquesr/users-test-golang/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID  primitive.ObjectID `json:"userId" bson:"userId,omitempty"`
	Name    string             `json:"name" bson:"name"`
	Address string             `json:"address" bson:"address"`
	Email   string             `json:"email" bson:"email"`
	Age     int32              `json:"age" bson:"age"`
}

type Handler struct {
	useCase user.UseCase
}

func NewHandler(useCase user.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) Create(c *gin.Context) {

	inp := &domain.User{
		Name:    "name",
		Email:   "email",
		Age:     20,
		Address: "address",
	}
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	login := c.MustGet(auth.CtxUserKey).(*domain.Login)

	if err := h.useCase.CreateLogin(c.Request.Context(), login, inp); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

type getResponse struct {
	Users []*User `json:"users"`
}

func (h *Handler) Get(c *gin.Context) {
	user := c.MustGet(auth.CtxUserKey).(*domain.Login)

	bms, err := h.useCase.GetLogin(c.Request.Context(), user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &getResponse{
		Users: toUsers(bms),
	})
}

type deleteInput struct {
	ID string `json:"id"`
}

func (h *Handler) Delete(c *gin.Context) {
	inp := new(deleteInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.MustGet(auth.CtxUserKey).(*domain.Login)

	if err := h.useCase.DeleteUser(c.Request.Context(), user, inp.ID); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func toUsers(bs []*domain.User) []*User {
	out := make([]*User, len(bs))

	for i, b := range bs {
		out[i] = toUser(b)
	}

	return out
}

func toUser(b *domain.User) *User {
	return &User{
		ID:      b.ID,
		UserID:  b.UserID,
		Name:    b.Name,
		Address: b.Address,
		Age:     b.Age,
		Email:   b.Email,
	}
}
