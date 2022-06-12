package transaction

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FormatterTransaction struct {
	ID          primitive.ObjectID `json:"id"`
	TrxType     string             `json:"trxType"`
	Amount      string             `json:"amount"`
	Description string             `json:"desc"`
	TrxMethod   string             `json:"method"`
	TrxID       string             `json:"trx_id"`
	RequestTime time.Time          `json:"created_at"`
}

func FormatTransaction(input Transaction) FormatterTransaction {
	format := FormatterTransaction{
		ID:          input.ID,
		TrxType:     input.TrxType,
		Amount:      input.Amount,
		Description: input.Desc,
		TrxMethod:   input.TrxMethod,
		TrxID:       input.TrxID,
		RequestTime: input.RequestTime,
	}
	return format
}

func FormatTransactions(input []Transaction) []FormatterTransaction {
	format := []FormatterTransaction{}
	for _, transaction := range input {
		format = append(format, FormatTransaction(transaction))
	}
	return format
}
