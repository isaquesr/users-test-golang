package auth

import (
	"context"

	"github.com/isaquesr/users-test-golang/domain"
)

const CtxUserKey = "login"

type UseCase interface {
	SignUp(ctx context.Context, username, password string) error
	SignIn(ctx context.Context, username, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*domain.Login, error)
}
