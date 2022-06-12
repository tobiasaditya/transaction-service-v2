package transaction

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	TrxType     string
	Amount      string
	Desc        string
	TrxMethod   string
	UserID      string
	TrxID       string
	RequestTime time.Time `bson:"requestTime,omitempty"`
}
