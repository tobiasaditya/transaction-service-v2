package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(user User) (User, error)
	GetUserByID(id primitive.ObjectID) (User, error)
	GetUserByPhone(phone string) (User, error)
}

type repository struct {
	db *mongo.Collection
}

func NewRepository(db *mongo.Collection) *repository {
	return &repository{db: db}
}

func (r *repository) Create(user User) (User, error) {
	_, err := r.db.InsertOne(context.TODO(), user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) GetUserByID(id primitive.ObjectID) (User, error) {
	filter := bson.D{{Key: "_id", Value: id}}
	foundUser := User{}
	err := r.db.FindOne(context.TODO(), filter).Decode(&foundUser)
	if err != nil {
		return foundUser, err
	}

	return foundUser, nil
}

func (r *repository) GetUserByPhone(phone string) (User, error) {
	filter := bson.D{{Key: "phoneNumber", Value: phone}}
	foundUser := User{}
	err := r.db.FindOne(context.TODO(), filter).Decode(&foundUser)
	if err != nil {
		return foundUser, err
	}

	return foundUser, nil
}

//================SQL=================
// import "gorm.io/gorm"

// type Repository interface {
// 	Create(user User) (User, error)
// 	GetUserByID(id int) (User, error)
// 	GetUserByPhone(phone string) (User, error)
// }

// type repository struct {
// 	db *gorm.DB
// }

// func NewRepository(db *gorm.DB) *repository {
// 	return &repository{db: db}
// }

// func (r *repository) Create(user User) (User, error) {
// 	err := r.db.Create(&user).Error
// 	if err != nil {
// 		return user, err
// 	}

// 	return user, nil
// }

// func (r *repository) GetUserByID(id int) (User, error) {
// 	foundUser := User{}
// 	err := r.db.Where("id = ?", id).Find(&foundUser).Error
// 	if err != nil {
// 		return foundUser, err
// 	}

// 	return foundUser, nil
// }

// func (r *repository) GetUserByPhone(phone string) (User, error) {
// 	foundUser := User{}
// 	err := r.db.Where("phone_number = ?", phone).Find(&foundUser).Error
// 	if err != nil {
// 		return foundUser, err
// 	}

// 	return foundUser, nil
// }
