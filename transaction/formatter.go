package transaction

import (
	"strconv"
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
	RequestTime time.Time          `json:"requestTime"`
}

type FormatterTransactionPage struct {
	TotalPurchase int                    `json:"total_purchase"`
	TotalIncome   int                    `json:"total_income"`
	TotalNet      int                    `json:"total_net"`
	NData         int                    `json:"n_data"`
	Content       []FormatterTransaction `json:"content"`
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

func FormatAccumulatorTransaction(input []Transaction) FormatterTransactionPage {
	format := FormatterTransactionPage{}
	format.Content = FormatTransactions(input)
	//Acumulate total data
	var totalPurchase int
	var totalIncome int

	for _, trx := range input {
		amount, _ := strconv.Atoi(trx.Amount)
		if trx.TrxType == "PURCHASE" {
			totalPurchase += amount
		} else if trx.TrxType == "INCOME" {
			totalIncome += amount
		}
	}
	totalNet := totalIncome - totalPurchase
	nData := len(input)
	format.TotalIncome = totalIncome
	format.TotalPurchase = totalPurchase
	format.TotalNet = totalNet
	format.NData = nData
	return format
}
