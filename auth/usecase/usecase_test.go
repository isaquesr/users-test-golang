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
		username = "user"
		password = "pass"

		ctx = context.Background()

		login = &domain.Login{
			Username: username,
			Password: "11f5639f22525155cb0b43573ee4212838c78d87", // sha1 of pass+salt
		}
	)

	// Sign Up
	repo.On("CreateLogin", login).Return(nil)
	err := uc.SignUp(ctx, username, password)
	assert.NoError(t, err)

	// Sign In (Get Auth Token)
	repo.On("GetLogin", login.Username, login.Password).Return(login, nil)
	token, err := uc.SignIn(ctx, username, password)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Verify token
	parsedUser, err := uc.ParseToken(ctx, token)
	assert.NoError(t, err)
	assert.Equal(t, login, parsedUser)
}
