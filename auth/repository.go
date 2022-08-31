package auth

import (
	"context"

	"github.com/isaquesr/users-test-golang/domain"
)

type LoginRepository interface {
	CreateLogin(ctx context.Context, login *domain.Login) error
	GetLogin(ctx context.Context, username, password string) (*domain.Login, error)
}
