package localstorage

import (
	"context"
	"testing"

	"github.com/isaquesr/users-test-golang/auth"
	"github.com/isaquesr/users-test-golang/domain"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetUser(t *testing.T) {
	s := NewUserLocalStorage()

	login := &domain.Login{
		ID:       primitive.NewObjectID(),
		Username: "user",
		Password: "password",
	}

	err := s.CreateLogin(context.Background(), login)
	assert.NoError(t, err)

	returnedUser, err := s.GetLogin(context.Background(), "user", "password")
	assert.NoError(t, err)
	assert.Equal(t, login, returnedUser)

	returnedUser, err = s.GetLogin(context.Background(), "user", "")
	assert.Error(t, err)
	assert.Equal(t, err, auth.ErrUserNotFound)
}
