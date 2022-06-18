package transaction

import (
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FormatterTransaction struct {
	ID          primitive.ObjectID `json:"id"`
	TrxType     string             `json:"trxType"`
	Amount      string             `json:"amount"`
	Description string             `json:"desc"`
	TrxMethod   string             `json:"trxMethod"`
	TrxID       string             `json:"trxId"`
	RequestTime string             `json:"requestTime"`
}

type FormatterTransactionPage struct {
	TotalPurchase int                    `json:"total_purchase"`
	TotalIncome   int                    `json:"total_income"`
	TotalNet      int                    `json:"total_net"`
	NData         int                    `json:"n_data"`
	Content       []FormatterTransaction `json:"content"`
}

type FormatterInvestmentPage struct {
	TotalInvestment int                    `json:"total_investment"`
	NData           int                    `json:"n_data"`
	Content         []FormatterTransaction `json:"content"`
}

func FormatTransaction(input Transaction) FormatterTransaction {
	format := FormatterTransaction{
		ID:          input.ID,
		TrxType:     input.TrxType,
		Amount:      input.Amount,
		Description: input.Desc,
		TrxMethod:   input.TrxMethod,
		TrxID:       input.TrxID,
		RequestTime: input.RequestTime.Format("2006-01-02 15:04:05"),
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

func FormatAccumulatorInvestment(input []Transaction) FormatterInvestmentPage {
	format := FormatterInvestmentPage{}
	format.Content = FormatTransactions(input)
	//Acumulate total data
	var totalInvestment int

	for _, trx := range input {
		amount, _ := strconv.Atoi(trx.Amount)
		totalInvestment += amount
	}
	nData := len(input)
	format.TotalInvestment = totalInvestment
	format.NData = nData
	return format
}
