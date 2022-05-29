package user

import "gorm.io/gorm"

type Repository interface {
	Create(user User) (User, error)
	GetUserByID(id int) (User, error)
	GetUserByPhone(phone string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(user User) (User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) GetUserByID(id int) (User, error) {
	foundUser := User{}
	err := r.db.Where("id = ?", id).Find(&foundUser).Error
	if err != nil {
		return foundUser, err
	}

	return foundUser, nil
}

func (r *repository) GetUserByPhone(phone string) (User, error) {
	foundUser := User{}
	err := r.db.Where("phone_number = ?", phone).Find(&foundUser).Error
	if err != nil {
		return foundUser, err
	}

	return foundUser, nil
}
