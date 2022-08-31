package mongo

import (
	"context"
	"errors"

	"github.com/isaquesr/users-test-golang/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Login struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
}

type LoginRepository struct {
	db *mongo.Collection
}

func NewLoginRepository(db *mongo.Database, collection string) *LoginRepository {
	return &LoginRepository{
		db: db.Collection(collection),
	}
}

func (r LoginRepository) CreateLogin(ctx context.Context, login *domain.Login) error {
	model := toMongoLogin(login)
	_, err := r.db.InsertOne(ctx, model)
	if IsDuplicate(err) {
		return domain.ErrUserAlreadyExists
	}

	return err
}

func (r LoginRepository) GetLogin(ctx context.Context, username, password string) (*domain.Login, error) {
	user := new(Login)
	err := r.db.FindOne(ctx, bson.M{
		"username": username,
		"password": password,
	}).Decode(user)

	if err != nil {
		return nil, err
	}

	return toModel(user), nil
}

func toMongoLogin(u *domain.Login) *Login {
	return &Login{
		Username: u.Username,
		Password: u.Password,
	}
}

func toModel(u *Login) *domain.Login {
	return &domain.Login{
		ID:       u.ID,
		Username: u.Username,
		Password: u.Password,
	}
}

func IsDuplicate(err error) bool {
	var e mongo.WriteException
	if errors.As(err, &e) {
		for _, we := range e.WriteErrors {
			if we.Code == 11000 {
				return true
			}
		}
	}

	return false
}
