package reporting

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DailyReport struct {
	ID          primitive.ObjectID `bson:"_id"`
	BodyWeight  float64            `bson:"bodyWeight"`
	Counter     int                `bson:"counter"`
	UserID      string             `bson:"userId"`
	RequestTime time.Time          `bson:"requestTime"`
}
