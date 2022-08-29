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

	user := &domain.User{
		ID:       primitive.NewObjectID(),
		Name:     "name",
		Password: "password",
		Email:    "email",
		Age:      20,
		Address:  "address",
	}

	err := s.CreateUser(context.Background(), user)
	assert.NoError(t, err)

	returnedUser, err := s.GetUser(context.Background(), "name", "password")
	assert.NoError(t, err)
	assert.Equal(t, user, returnedUser)

	returnedUser, err = s.GetUser(context.Background(), "name", "")
	assert.Error(t, err)
	assert.Equal(t, err, auth.ErrUserNotFound)
}
