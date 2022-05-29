package transaction

import "time"

type Transaction struct {
	ID            int
	Type          string
	Amount        int
	Description   string
	Method        string
	UserID        int
	TransactionID string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
