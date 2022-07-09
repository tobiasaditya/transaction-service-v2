package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id, omitempty"`
	FullName    string             `bson:"fullName"`
	Email       string             `bson:"email"`
	PhoneNumber string             `bson:"phoneNumber"`
	Token       string             `bson:"token"`
	Password    string             `bson:"password"`
	CreateTime  time.Time          `bson:"createTime,omitempty"`
}
