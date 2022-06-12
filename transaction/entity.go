package transaction

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID          primitive.ObjectID `bson:"_id"`
	TrxType     string             `bson:"trxType"`
	Amount      string             `bson:"amount"`
	Desc        string             `bson:"desc"`
	TrxMethod   string             `bson:"trxMethod"`
	UserID      string             `bson:"userId"`
	TrxID       string             `bson:"trxId"`
	RequestTime time.Time          `bson:"requestTime"`
}
