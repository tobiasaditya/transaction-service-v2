package transaction

import "time"

type Service interface {
	CreateTransaction(input InputTransaction) (Transaction, error)
	GetTransactions(userID string, start time.Time, end time.Time) ([]Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s service) CreateTransaction(input InputTransaction) (Transaction, error) {
	transaction := Transaction{
		TrxType:   input.TrxType,
		Amount:    input.Amount,
		Desc:      input.Desc,
		TrxMethod: input.Method,
		UserID:    "6189f1796bb08e7dc15fe3ef",
		TrxID:     "TRX-INV-20201921",
	}

	transaction, err := s.repository.Create(transaction)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (s service) GetTransactions(userID string, start time.Time, end time.Time) ([]Transaction, error) {
	transactions := []Transaction{}

	transactions, err := s.repository.GetTransactionsByUserID(userID, start, end)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
