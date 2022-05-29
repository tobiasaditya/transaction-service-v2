package user

import "time"

type User struct {
	ID          int
	FullName    string
	Email       string
	PhoneNumber string
	CreatedAt   time.Time
	Password    string
}
