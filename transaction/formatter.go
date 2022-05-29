package transaction

type FormatterTransaction struct {
	ID            int    `json:"id"`
	Type          string `json:"trxType"`
	Amount        int    `json:"amount"`
	Description   string `json:"desc"`
	Method        string `json:"method"`
	TransactionID string `json:"trx_id"`
}

func FormatTransaction(input Transaction) FormatterTransaction {
	format := FormatterTransaction{
		ID:            input.ID,
		Type:          input.Type,
		Amount:        input.Amount,
		Description:   input.Description,
		Method:        input.Method,
		TransactionID: input.TransactionID,
	}
	return format
}
