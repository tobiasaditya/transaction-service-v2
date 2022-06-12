package transaction

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	CreateTransaction(input InputTransaction, userID string) (Transaction, error)
	GetTransactions(userID string, start time.Time, end time.Time) ([]Transaction, error)
	GetInvestments(userID string) ([]Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s service) CreateTransaction(input InputTransaction, userID string) (Transaction, error) {
	transaction := Transaction{
		ID:          primitive.NewObjectID(),
		TrxType:     input.TrxType,
		Amount:      input.Amount,
		Desc:        input.Desc,
		TrxMethod:   input.Method,
		UserID:      userID,
		TrxID:       uuid.New().String(),
		RequestTime: time.Now(),
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

func (s service) GetInvestments(userID string) ([]Transaction, error) {
	transactions := []Transaction{}

	transactions, err := s.repository.GetInvestmentsByUserID(userID)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
