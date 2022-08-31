package localstorage

import (
	"context"
	"sync"

	"github.com/isaquesr/users-test-golang/auth"
	"github.com/isaquesr/users-test-golang/domain"
)

type UserLocalStorage struct {
	users map[string]*domain.Login
	mutex *sync.Mutex
}

func NewUserLocalStorage() *UserLocalStorage {
	return &UserLocalStorage{
		users: make(map[string]*domain.Login),
		mutex: new(sync.Mutex),
	}
}

func (s *UserLocalStorage) CreateLogin(ctx context.Context, user *domain.Login) error {
	s.mutex.Lock()
	s.users[user.ID.String()] = user
	s.mutex.Unlock()

	return nil
}

func (s *UserLocalStorage) GetLogin(ctx context.Context, username, password string) (*domain.Login, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, user := range s.users {
		if user.Username == username && user.Password == password {
			return user, nil
		}
	}

	return nil, auth.ErrUserNotFound
}
