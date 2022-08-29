package auth

import (
	"context"

	"github.com/isaquesr/users-test-golang/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUser(ctx context.Context, name, password string) (*domain.User, error)
}
