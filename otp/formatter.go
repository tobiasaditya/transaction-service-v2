package otp

import "go.mongodb.org/mongo-driver/bson/primitive"

type FormatterOtp struct {
	ID primitive.ObjectID `json:"id"`
}

func FormatOtp(otp Otp) FormatterOtp {
	format := FormatterOtp{
		ID: otp.ID,
	}
	return format
}
