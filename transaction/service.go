package transaction

type Service interface {
	CreateTransaction(input InputTransaction) (Transaction, error)
	GetTransactions(userID int) ([]Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s service) CreateTransaction(input InputTransaction) (Transaction, error) {
	transaction := Transaction{
		Type:          input.Type,
		Amount:        input.Amount,
		Description:   input.Description,
		Method:        input.Method,
		UserID:        1,
		TransactionID: "TRX-INV-20201921",
	}

	transaction, err := s.repository.Create(transaction)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (s service) GetTransactions(userID int) ([]Transaction, error) {
	transactions := []Transaction{}

	transactions, err := s.repository.GetTransactionsByUserID(userID)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
