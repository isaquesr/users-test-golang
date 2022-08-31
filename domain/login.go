package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Login struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
}
