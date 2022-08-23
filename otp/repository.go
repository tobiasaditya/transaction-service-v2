package otp

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(otp Otp) (Otp, error)
	Update(id primitive.ObjectID, isVerified bool) (bool, error)
	Find(id primitive.ObjectID) (Otp, error)
}

type repository struct {
	db *mongo.Collection
}

func NewRepository(db *mongo.Collection) *repository {
	return &repository{db: db}
}

func (r *repository) Create(otp Otp) (Otp, error) {
	_, err := r.db.InsertOne(context.TODO(), otp)
	if err != nil {
		return otp, err
	}

	return otp, nil
}

func (r *repository) Update(id primitive.ObjectID, isVerified bool) (bool, error) {
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "otpVerified", Value: isVerified}}}}
	_, err := r.db.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *repository) Find(id primitive.ObjectID) (Otp, error) {
	filter := bson.D{{Key: "_id", Value: id}}
	foundOtp := Otp{}
	err := r.db.FindOne(context.TODO(), filter).Decode(&foundOtp)
	if err != nil {
		return foundOtp, err
	}

	return foundOtp, nil
}
