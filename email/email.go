package email

import (
	"fmt"
	"net/smtp"
	"os"
	"transaction-service-v2/otp"
)

func SendEmailOtp(otp otp.Otp) error {

	from := os.Getenv("G_EMAIL")
	password := os.Getenv("G_PASSWORD")

	toEmailAddress := otp.Receiver
	to := []string{toEmailAddress}

	host := "smtp.gmail.com" //os.Getenv("G_HOST")
	port := "587"            //os.Getenv("G_PORT")
	address := host + ":" + port

	subject := "Subject: Registration Obider Personal App\n"
	body := fmt.Sprintf("This is you One Time Password (OTP) %s", otp.Value)
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		return err
	}
	return nil

}
