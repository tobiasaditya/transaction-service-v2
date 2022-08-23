package otp

import (
	"crypto/rand"
	"errors"
	"io"
	"time"
	"transaction-service-v2/util"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

type Service interface {
	CreateOtp(email string) (Otp, error)
	VerifyOtp(id string, value string) (Otp, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateOtp(email string) (Otp, error) {
	newOtp := Otp{
		ID:         primitive.NewObjectID(),
		Value:      GenerateNumber(6),
		Expired:    util.CTimeNow().Add(time.Minute * 1),
		Receiver:   email,
		IsVerified: false,
	}

	newOtp, err := s.repository.Create(newOtp)

	if err != nil {
		return newOtp, err
	}
	return newOtp, nil
}

func (s *service) VerifyOtp(id string, value string) (Otp, error) {
	newID, _ := primitive.ObjectIDFromHex(id)
	foundOtp, err := s.repository.Find(newID)
	if err != nil {
		return foundOtp, err
	}

	//Check expired date
	timeNow := util.CTimeNow()
	if timeNow.After(foundOtp.Expired) || foundOtp.IsVerified {
		return foundOtp, errors.New("expired otp")
	}

	//Check value
	if value != foundOtp.Value {
		return foundOtp, errors.New("incorrect otp")
	}

	//Update otp status
	_, err = s.repository.Update(newID, true)
	if err != nil {
		return foundOtp, err
	}
	return foundOtp, nil
}

func GenerateNumber(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}
