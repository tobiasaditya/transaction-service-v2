package otp

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Otp struct {
	ID         primitive.ObjectID `bson:"_id, omitempty"`
	Value      string             `bson:"otpValue"`
	Expired    time.Time          `bson:"otpExpired"`
	Receiver   string             `bson:"otpReceiver"`
	IsVerified bool               `bson:"otpVerified"`
}
