package transaction

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Create(transaction Transaction) (Transaction, error)
	GetTransactionsByUserID(userID string, start time.Time, end time.Time) ([]Transaction, error)
	GetInvestmentsByUserID(userID string) ([]Transaction, error)
}

type repository struct {
	db *mongo.Collection
}

func NewRepository(db *mongo.Collection) *repository {
	return &repository{db: db}
}

func (r *repository) Create(transaction Transaction) (Transaction, error) {
	_, err := r.db.InsertOne(context.TODO(), transaction)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repository) GetTransactionsByUserID(userID string, start time.Time, end time.Time) ([]Transaction, error) {
	trxType := []string{"PURCHASE", "INCOME"}
	filter := bson.D{
		{Key: "userId", Value: userID},
		{Key: "requestTime", Value: bson.D{
			{Key: "$gte", Value: start},
			{Key: "$lte", Value: end},
		}},
		{Key: "trxType", Value: bson.D{
			{Key: "$in", Value: trxType},
		}},
	}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "requestTime", Value: -1}})
	// fmt.Println(filter)
	transactions := []Transaction{}
	cursor, err := r.db.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return transactions, err
	}
	err = cursor.All(context.TODO(), &transactions)
	if err != nil {
		fmt.Println(err)
		return transactions, err
	}
	return transactions, nil
}

func (r *repository) GetInvestmentsByUserID(userID string) ([]Transaction, error) {
	filter := bson.D{
		{Key: "userId", Value: userID},
		{Key: "trxType", Value: "INVESTMENT"},
	}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "_id", Value: -1}})
	// fmt.Println(filter)
	transactions := []Transaction{}
	cursor, err := r.db.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return transactions, err
	}
	err = cursor.All(context.TODO(), &transactions)
	if err != nil {
		fmt.Println(err)
		return transactions, err
	}
	return transactions, nil
}

//========SQL============
// type repository struct {
// 	db *gorm.DB
// }

// func NewRepository(db *gorm.DB) *repository {
// 	return &repository{db: db}
// }

// func (r *repository) Create(transaction Transaction) (Transaction, error) {
// 	err := r.db.Create(&transaction).Error
// 	if err != nil {
// 		return transaction, err
// 	}
// 	return transaction, nil
// }

// func (r *repository) GetTransactionsByUserID(userID int) ([]Transaction, error) {
// 	transactions := []Transaction{}
// 	err := r.db.Where("user_id = ?", userID).Find(&transactions).Error
// 	if err != nil {
// 		return transactions, err
// 	}
// 	return transactions, nil
// }
