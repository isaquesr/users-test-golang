package usecase

import (
	"context"

	"github.com/isaquesr/users-test-golang/domain"
	"github.com/isaquesr/users-test-golang/user"
)

type UserUseCase struct {
	userRepo user.Repository
}

func NewUserUseCase(userRepo user.Repository) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

func (b UserUseCase) CreateLogin(ctx context.Context, login *domain.Login, user *domain.User) error {
	usr := &domain.User{
		Address: user.Address,
		Email:   user.Email,
		Age:     user.Age,
		Name:    user.Name,
	}

	return b.userRepo.CreateLogin(ctx, login, usr)
}

func (b UserUseCase) GetLogin(ctx context.Context, login *domain.Login) ([]*domain.User, error) {
	return b.userRepo.GetLogin(ctx, login)
}

func (b UserUseCase) DeleteUser(ctx context.Context, login *domain.Login, id string) error {
	return b.userRepo.DeleteUser(ctx, login, id)
}
