package mongo

import (
	"context"
	"errors"

	"github.com/isaquesr/users-test-golang/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDto struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Age      int32              `json:"age" bson:"age"`
	Password string             `json:"password" bson:"password"`
	Address  string             `json:"address" bson:"address"`
}

type UserRepository struct {
	db *mongo.Collection
}

func NewUserRepository(db *mongo.Database, collection string) *UserRepository {
	return &UserRepository{
		db: db.Collection(collection),
	}
}

func (r *UserRepository) Create(ctx context.Context, user domain.User) error {
	users := toMongoUser(user)
	_, err := r.db.InsertOne(ctx, users)
	if IsDuplicate(err) {
		return domain.ErrUserAlreadyExists
	}

	return err
}

func (r UserRepository) GetUser(ctx context.Context, email, password string) (*domain.User, error) {
	user := new(UserDto)
	err := r.db.FindOne(ctx, bson.M{
		"email":    email,
		"password": password,
	}).Decode(user)

	if err != nil {
		return nil, err
	}

	return toModel(user), nil
}

func toModel(u *UserDto) *domain.User {
	return &domain.User{
		ID:       u.ID,
		Name:     u.Name,
		Address:  u.Address,
		Age:      u.Age,
		Email:    u.Email,
		Password: u.Password,
	}
}

func toMongoUser(user domain.User) *UserDto {
	return &UserDto{
		ID:       user.ID,
		Name:     user.Name,
		Address:  user.Address,
		Age:      user.Age,
		Email:    user.Email,
		Password: user.Password,
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
