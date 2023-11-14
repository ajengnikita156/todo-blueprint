package helpers

import (
	"fmt"

	"gopkg.in/gomail.v2"

)

func SendResetPasswordEmail(toEmail string, token string) error {
	resetLink := fmt.Sprintf("https://localhost:7080/reset-password?token=%s", token)

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "ajengnikita14@gmail.com")
	mailer.SetHeader("To", toEmail)
	mailer.SetHeader("Subject", "Reset Password")
	mailer.SetBody("text/html", fmt.Sprintf("Klik <a href='%s'>di sini</a> untuk mereset password Anda.", resetLink))

	dialer := gomail.NewDialer("smtp.gmail.com", 587, "ajengnikita14@gmail.com", "fird jdwa rujm xlyq")

	if err := dialer.DialAndSend(mailer); err != nil {
		return err
	}

	return nil
}