package user

import "errors"

type Service interface {
	CreateUser(input InputUser) (User, error)
	FindUserByID(id int) (User, error)
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
		FullName:    input.FullName,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Password:    input.Password,
	}

	user, err := s.repository.Create(user)

	if err != nil {
		return user, err
	}
	return user, nil
}

func (s service) FindUserByID(id int) (User, error) {
	foundUser, err := s.repository.GetUserByID(id)
	if err != nil {
		return foundUser, err
	}
	return foundUser, nil
}

func (s service) Login(input InputLogin) (User, error) {
	foundUser, err := s.repository.GetUserByPhone(input.PhoneNumber)
	if err != nil {
		return foundUser, err
	}

	if foundUser.ID == 0 {
		return User{}, errors.New("User not found")
	}

	//Jika nemu, cek password
	if foundUser.Password != input.Password {
		return foundUser, errors.New("Incorrect password")
	}

	return foundUser, nil
}
