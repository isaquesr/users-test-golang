package mock

import (
	"context"

	"github.com/isaquesr/users-test-golang/domain"
	"github.com/stretchr/testify/mock"
)

type UserStorageMock struct {
	mock.Mock
}

func (s *UserStorageMock) CreateLogin(ctx context.Context, user *domain.Login) error {
	args := s.Called(user)

	return args.Error(0)
}

func (s *UserStorageMock) GetLogin(ctx context.Context, username, password string) (*domain.Login, error) {
	args := s.Called(username, password)

	return args.Get(0).(*domain.Login), args.Error(1)
}
