package localcache

import (
	"context"
	"sync"

	"github.com/isaquesr/users-test-golang/domain"

	"github.com/isaquesr/users-test-golang/user"
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

func (s *UserLocalStorage) CreateLogin(ctx context.Context, login *domain.Login, user *domain.User) error {
	user.UserID = login.ID

	s.mutex.Lock()
	s.users[user.ID.String()] = user
	s.mutex.Unlock()

	return nil
}

func (s *UserLocalStorage) GetLogin(ctx context.Context, login *domain.Login) ([]*domain.User, error) {
	users := make([]*domain.User, 0)

	s.mutex.Lock()
	for _, bm := range s.users {
		if bm.UserID == login.ID {
			users = append(users, bm)
		}
	}
	s.mutex.Unlock()

	return users, nil
}

func (s *UserLocalStorage) DeleteUser(ctx context.Context, login *domain.Login, id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	bm, ex := s.users[id]
	if ex && bm.UserID == login.ID {
		delete(s.users, id)
		return nil
	}

	return user.ErrUserNotFound
}
