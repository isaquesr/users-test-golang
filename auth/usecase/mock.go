package usecase

import (
	"context"

	"github.com/isaquesr/users-test-golang/domain"
	"github.com/stretchr/testify/mock"
)

type AuthUseCaseMock struct {
	mock.Mock
}

func (m *AuthUseCaseMock) SignUp(ctx context.Context, u *domain.User) error {
	args := m.Called(u)

	return args.Error(0)
}

func (m *AuthUseCaseMock) SignIn(ctx context.Context, username, password string) (string, error) {
	args := m.Called(username, password)

	return args.Get(0).(string), args.Error(1)
}

func (m *AuthUseCaseMock) ParseToken(ctx context.Context, accessToken string) (*domain.User, error) {
	args := m.Called(accessToken)

	return args.Get(0).(*domain.User), args.Error(1)
}
