package user

import (
	"context"

	"github.com/isaquesr/users-test-golang/domain"
)

type UseCase interface {
	CreateLogin(ctx context.Context, login *domain.Login, user *domain.User) error
	GetLogin(ctx context.Context, login *domain.Login) ([]*domain.User, error)
	DeleteUser(ctx context.Context, login *domain.Login, id string) error
}
