package email

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmail(email string) {

	from := os.Getenv("G_EMAIL")
	password := os.Getenv("G_PASSWORD")
	fmt.Println(from)
	fmt.Println(password)
	toEmailAddress := email
	to := []string{toEmailAddress}

	host := "smtp.gmail.com" //os.Getenv("G_HOST")
	port := "587"            //os.Getenv("G_PORT")
	address := host + ":" + port

	subject := "Subject: Registration Obider Personal App\n"
	body := "This is the body of the mail"
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		panic(err)
	}

}
