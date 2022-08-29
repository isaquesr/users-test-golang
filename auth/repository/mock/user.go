package mock

import (
	"context"

	"github.com/isaquesr/users-test-golang/domain"
	"github.com/stretchr/testify/mock"
)

type UserStorageMock struct {
	mock.Mock
}

func (s *UserStorageMock) CreateUser(ctx context.Context, user *domain.User) error {
	args := s.Called(user)

	return args.Error(0)
}

func (s *UserStorageMock) GetUser(ctx context.Context, name, password string) (*domain.User, error) {
	args := s.Called(name, password)

	return args.Get(0).(*domain.User), args.Error(1)
}
