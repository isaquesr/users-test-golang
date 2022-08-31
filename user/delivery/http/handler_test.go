package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/isaquesr/users-test-golang/auth"
	"github.com/isaquesr/users-test-golang/domain"
	"github.com/isaquesr/users-test-golang/user/usecase"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreate(t *testing.T) {
	testUser := &domain.Login{
		Username: "testuser",
		Password: "testpass",
	}

	r := gin.Default()
	group := r.Group("/api", func(c *gin.Context) {
		c.Set(auth.CtxUserKey, testUser)
	})

	uc := new(usecase.UserUseCaseMock)

	RegisterHTTPEndpoints(group, uc)

	inp := &domain.User{
		Name:    "name",
		Email:   "email",
		Age:     20,
		Address: "address",
	}

	body, err := json.Marshal(inp)
	assert.NoError(t, err)

	uc.On("CreateLogin", testUser, inp).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGet(t *testing.T) {
	testUser := &domain.Login{
		Username: "testuser",
		Password: "testpass",
	}

	r := gin.Default()
	group := r.Group("/api", func(c *gin.Context) {
		c.Set(auth.CtxUserKey, testUser)
	})

	uc := new(usecase.UserUseCaseMock)

	RegisterHTTPEndpoints(group, uc)

	bms := make([]*domain.User, 5)
	for i := 0; i < 5; i++ {
		bms[i] = &domain.User{
			ID:      primitive.NewObjectID(),
			UserID:  primitive.NewObjectID(),
			Name:    "name",
			Email:   "email",
			Age:     20,
			Address: "address",
		}
	}

	uc.On("GetLogin", testUser).Return(bms, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/users", nil)
	r.ServeHTTP(w, req)

	expectedOut := &getResponse{Users: toUsers(bms)}

	expectedOutBody, err := json.Marshal(expectedOut)
	assert.NoError(t, err)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(expectedOutBody), w.Body.String())
}

func TestDelete(t *testing.T) {
	testUser := &domain.Login{
		Username: "testuser",
		Password: "testpass",
	}

	r := gin.Default()
	group := r.Group("/api", func(c *gin.Context) {
		c.Set(auth.CtxUserKey, testUser)
	})

	uc := new(usecase.UserUseCaseMock)

	RegisterHTTPEndpoints(group, uc)

	inp := &deleteInput{
		ID: "id",
	}

	body, err := json.Marshal(inp)
	assert.NoError(t, err)

	uc.On("DeleteUser", testUser, inp.ID).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/users", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
