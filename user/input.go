package user

type InputUser struct {
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type InputLogin struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
