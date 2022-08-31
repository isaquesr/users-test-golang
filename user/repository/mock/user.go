package mock

import (
	"context"

	"github.com/isaquesr/users-test-golang/domain"
	"github.com/stretchr/testify/mock"
)

type UserStorageMock struct {
	mock.Mock
}

func (s *UserStorageMock) CreateLogin(ctx context.Context, login *domain.Login, user *domain.User) error {
	args := s.Called(login, user)

	return args.Error(0)
}

func (s *UserStorageMock) GetUsers(ctx context.Context, login *domain.Login) ([]*domain.User, error) {
	args := s.Called(login)

	return args.Get(0).([]*domain.User), args.Error(1)
}

func (s *UserStorageMock) DeleteUser(ctx context.Context, login *domain.Login, id string) error {
	args := s.Called(login, id)

	return args.Error(0)
}
