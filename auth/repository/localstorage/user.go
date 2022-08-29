package localstorage

import (
	"context"
	"sync"

	"github.com/isaquesr/users-test-golang/auth"
	"github.com/isaquesr/users-test-golang/domain"
)

type UserLocalStorage struct {
	users map[string]*domain.User
	mutex *sync.Mutex
}

func NewUserLocalStorage() *UserLocalStorage {
	return &UserLocalStorage{
		users: make(map[string]*domain.User),
		mutex: new(sync.Mutex),
	}
}

func (s *UserLocalStorage) CreateUser(ctx context.Context, user *domain.User) error {
	s.mutex.Lock()
	s.users[user.ID.String()] = user
	s.mutex.Unlock()

	return nil
}

func (s *UserLocalStorage) GetUser(ctx context.Context, name, password string) (*domain.User, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, user := range s.users {
		if user.Name == name && user.Password == password {
			return user, nil
		}
	}

	return nil, auth.ErrUserNotFound
}
