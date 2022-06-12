package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type FormatterUser struct {
	ID       primitive.ObjectID `json:"id"`
	FullName string             `json:"full_name"`
	Email    string             `json:"email"`
}

func FormatUser(user User) FormatterUser {
	format := FormatterUser{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
	}
	return format
}
