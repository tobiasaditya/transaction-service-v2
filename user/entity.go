package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id, omitempty"`
	FullName    string
	Email       string
	PhoneNumber string
	Token       string
	Password    string
	CreateTime  time.Time `bson:"createTime,omitempty"`
}
