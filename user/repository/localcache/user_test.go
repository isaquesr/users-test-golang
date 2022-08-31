package localcache

import (
	"context"
	"testing"

	"github.com/isaquesr/users-test-golang/domain"
	"github.com/isaquesr/users-test-golang/user"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetUsers(t *testing.T) {
	login := &domain.Login{ID: primitive.NewObjectID()}

	s := NewUserLocalStorage()

	for i := 0; i < 10; i++ {
		usr := &domain.User{
			ID:     primitive.NewObjectID(),
			UserID: login.ID,
		}

		err := s.CreateLogin(context.Background(), login, usr)
		assert.NoError(t, err)
	}

	returnedUsers, err := s.GetLogin(context.Background(), login)
	assert.NoError(t, err)

	assert.Equal(t, 10, len(returnedUsers))
}

func TestDeleteUser(t *testing.T) {
	id1 := primitive.NewObjectID()
	s1 := id1.String()
	hex1 := s1[10:34]
	id2, err := primitive.ObjectIDFromHex(hex1)

	id22 := primitive.NewObjectID()
	s22 := id22.String()
	hex22 := s22[10:34]
	id3, err := primitive.ObjectIDFromHex(hex22)

	idbm := primitive.NewObjectID()
	sbm := idbm.String()
	hexbm := sbm[10:34]
	bmID, err := primitive.ObjectIDFromHex(hexbm)

	login1 := &domain.Login{ID: id3}
	login2 := &domain.Login{ID: id2}

	bm := &domain.User{ID: bmID}

	s := NewUserLocalStorage()

	err = s.CreateLogin(context.Background(), login1, bm)
	assert.NoError(t, err)

	err = s.DeleteUser(context.Background(), login1, bmID.String())
	assert.NoError(t, err)

	err = s.CreateLogin(context.Background(), login1, bm)
	assert.NoError(t, err)

	err = s.DeleteUser(context.Background(), login2, bmID.String())
	assert.Error(t, err)
	assert.Equal(t, err, user.ErrUserNotFound)
}
