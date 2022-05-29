package transaction

import "time"

type FormatterTransaction struct {
	ID            int       `json:"id"`
	Type          string    `json:"trxType"`
	Amount        int       `json:"amount"`
	Description   string    `json:"desc"`
	Method        string    `json:"method"`
	TransactionID string    `json:"trx_id"`
	CreatedAt     time.Time `json:"created_at"`
}

func FormatTransaction(input Transaction) FormatterTransaction {
	format := FormatterTransaction{
		ID:            input.ID,
		Type:          input.Type,
		Amount:        input.Amount,
		Description:   input.Description,
		Method:        input.Method,
		TransactionID: input.TransactionID,
		CreatedAt:     input.CreatedAt,
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
