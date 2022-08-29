package usecase

import (
	"context"
	"testing"

	"github.com/isaquesr/users-test-golang/auth/repository/mock"
	"github.com/isaquesr/users-test-golang/domain"
	"github.com/stretchr/testify/assert"
)

func TestAuthFlow(t *testing.T) {
	repo := new(mock.UserStorageMock)

	uc := NewAuthUseCase(repo, "salt", []byte("secret"), 86400)

	var (
		name     = "name"
		password = "pass"
		usr      = &domain.User{
			Name:     name,
			Password: password,
			Address:  "Rua teste, 40",
			Age:      20,
			Email:    "teste@teste.com.br",
		}

		ctx = context.Background()

		user = &domain.User{
			Name:     name,
			Password: "11f5639f22525155cb0b43573ee4212838c78d87", // sha1 of pass+salt
			Address:  "Rua teste, 40",
			Age:      20,
			Email:    "teste@teste.com.br",
		}
	)

	// Sign Up
	repo.On("CreateUser", user).Return(nil)
	err := uc.SignUp(ctx, usr)
	assert.NoError(t, err)

	// Sign In (Get Auth Token)
	repo.On("GetUser", user.Name, user.Password).Return(user, nil)
	token, err := uc.SignIn(ctx, name, password)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Verify token
	parsedUser, err := uc.ParseToken(ctx, token)
	assert.NoError(t, err)
	assert.Equal(t, user, parsedUser)
}
