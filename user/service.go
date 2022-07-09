package user

import (
	"errors"
	"transaction-service-v2/util"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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

	//Hashed password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)
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
			return foundUser, errors.New("Incorrect username/password")
		}
		return foundUser, err
	}

	//Jika nemu, cek password
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(input.Password))
	if err != nil {
		return foundUser, errors.New("Incorrect username/password")
	}

	return foundUser, nil
}
