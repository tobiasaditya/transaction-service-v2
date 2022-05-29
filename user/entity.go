package user

import "time"

type User struct {
	ID          int
	FullName    string
	Email       string
	PhoneNumber string
	Token       string
	Password    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
