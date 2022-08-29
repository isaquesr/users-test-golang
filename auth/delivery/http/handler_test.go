package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/isaquesr/users-test-golang/auth/usecase"
	"github.com/isaquesr/users-test-golang/domain"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestSignUp(t *testing.T) {
	r := gin.Default()
	uc := new(usecase.AuthUseCaseMock)

	RegisterHTTPEndpoints(r, uc)

	user := &domain.User{
		ID:       primitive.NewObjectID(),
		Name:     "name",
		Password: "password",
		Email:    "email",
		Age:      20,
		Address:  "address",
	}

	body, err := json.Marshal(user)
	assert.NoError(t, err)

	uc.On("SignUp", user).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/auth/sign-up", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestSignIn(t *testing.T) {
	r := gin.Default()
	uc := new(usecase.AuthUseCaseMock)

	RegisterHTTPEndpoints(r, uc)

	signUpBody := &SignIn{
		Name:     "testuser",
		Password: "testpass",
	}

	body, err := json.Marshal(signUpBody)
	assert.NoError(t, err)

	uc.On("SignIn", signUpBody.Name, signUpBody.Password).Return("jwt", nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/auth/sign-in", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"token\":\"jwt\"}", w.Body.String())
}
