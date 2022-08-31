package mongo

import (
	"context"
	"errors"

	"github.com/isaquesr/users-test-golang/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID  primitive.ObjectID `json:"userId" bson:"userId,omitempty"`
	Name    string             `json:"name" bson:"name"`
	Address string             `json:"address" bson:"address"`
	Email   string             `json:"email" bson:"email"`
	Age     int32              `json:"age" bson:"age"`
}

type UserRepository struct {
	db *mongo.Collection
}

func NewUserRepository(db *mongo.Database, collection string) *UserRepository {
	return &UserRepository{
		db: db.Collection(collection),
	}
}

func (r UserRepository) CreateLogin(ctx context.Context, login *domain.Login, user *domain.User) error {
	user.UserID = login.ID

	model := toModel(user)

	_, err := r.db.InsertOne(ctx, model)
	if IsDuplicate(err) {
		return domain.ErrUserAlreadyExists
	}

	return err
}

func (r UserRepository) GetLogin(ctx context.Context, login *domain.Login) ([]*domain.User, error) {
	uid, _ := primitive.ObjectIDFromHex(login.ID.String())
	cur, err := r.db.Find(ctx, bson.M{
		"userId": uid,
	})

	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	out := make([]*User, 0)

	for cur.Next(ctx) {
		user := new(User)
		err := cur.Decode(user)
		if err != nil {
			return nil, err
		}

		out = append(out, user)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return toUsers(out), nil
}

func (r UserRepository) DeleteUser(ctx context.Context, login *domain.Login, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	uID, _ := primitive.ObjectIDFromHex(login.ID.String())

	_, err := r.db.DeleteOne(ctx, bson.M{"_id": objID, "userId": uID})
	return err
}

func toModel(b *domain.User) *User {
	return &User{
		UserID:  b.UserID,
		Name:    b.Name,
		Address: b.Address,
		Email:   b.Email,
		Age:     b.Age,
	}
}

func toUser(b *User) *domain.User {
	return &domain.User{
		ID:      b.ID,
		UserID:  b.UserID,
		Name:    b.Name,
		Address: b.Address,
		Email:   b.Email,
		Age:     b.Age,
	}
}

func toUsers(bs []*User) []*domain.User {
	out := make([]*domain.User, len(bs))

	for i, b := range bs {
		out[i] = toUser(b)
	}

	return out
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
