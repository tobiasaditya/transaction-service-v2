package user

type FormatterUser struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

func FormatUser(user User) FormatterUser {
	format := FormatterUser{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
	}
	return format
}
