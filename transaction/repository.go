package transaction

import "gorm.io/gorm"

type Repository interface {
	Create(transaction Transaction) (Transaction, error)
	GetTransactionsByUserID(userID int) ([]Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repository) GetTransactionsByUserID(userID int) ([]Transaction, error) {
	transactions := []Transaction{}
	err := r.db.Where("user_id = ?", userID).Find(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
