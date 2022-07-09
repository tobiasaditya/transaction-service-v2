package user

import (
	"errors"
	"transaction-service-v2/util"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	CreateUser(input InputUser) (User, error)
	FindUserByID(id string) (User, error)
	Login(input InputLogin) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s service) CreateUser(input InputUser) (User, error) {
	user := User{
		ID:          primitive.NewObjectID(),
		FullName:    input.FullName,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Password:    input.Password,
		CreateTime:  util.CTimeNow(),
	}

	//Check for duplicate phonenumber(username)
	dupUser, err := s.repository.GetUserByPhone(user.PhoneNumber)
	if dupUser.PhoneNumber == user.PhoneNumber {
		return dupUser, errors.New("Username has been used")
	}

	user, err = s.repository.Create(user)

	if err != nil {
		return user, err
	}
	return user, nil
}

func (s service) FindUserByID(id string) (User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return User{}, err
	}
	foundUser, err := s.repository.GetUserByID(objectID)
	if err != nil {
		return foundUser, err
	}
	return foundUser, nil
}

func (s service) Login(input InputLogin) (User, error) {
	foundUser, err := s.repository.GetUserByPhone(input.Username)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return foundUser, errors.New("User not found")
		}
		return foundUser, err
	}

	//Jika nemu, cek password
	if foundUser.Password != input.Password {
		return foundUser, errors.New("Incorrect password")
	}

	return foundUser, nil
}
