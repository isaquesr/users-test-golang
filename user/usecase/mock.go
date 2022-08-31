package usecase

import (
	"context"

	"github.com/isaquesr/users-test-golang/domain"
	"github.com/stretchr/testify/mock"
)

type UserUseCaseMock struct {
	mock.Mock
}

func (m UserUseCaseMock) CreateLogin(ctx context.Context, login *domain.Login, user *domain.User) error {
	args := m.Called(login, user)

	return args.Error(0)
}

func (m UserUseCaseMock) GetLogin(ctx context.Context, login *domain.Login) ([]*domain.User, error) {
	args := m.Called(login)

	return args.Get(0).([]*domain.User), args.Error(1)
}

func (m UserUseCaseMock) DeleteUser(ctx context.Context, login *domain.Login, id string) error {
	args := m.Called(login, id)

	return args.Error(0)
}
