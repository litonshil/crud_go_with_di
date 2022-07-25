package utils

import (
	"fmt"
	"net/smtp"

	"github.com/litonshil/crud_go_echo/config"
	"github.com/litonshil/crud_go_echo/pkg/models"
)

// SendEmail sends an email to user
func SendEmail(user *models.User) error {
	to := []string{user.Email}

	address := config.GetConfig().SmtpHost + ":" + config.GetConfig().SmtpPort

	subject := "Welcome to user management"
	body := fmt.Sprintf("Your credentials for login are given below: \n")
	body += fmt.Sprintf("Username: %s \n", user.Email)
	body += fmt.Sprintf("Password: %s \n", user.Password)

	message := fmt.Sprintf("From: %s\r\n", config.GetConfig().Email)
	message += fmt.Sprintf("To: %s\r\n", to)
	message += fmt.Sprintf("Subject: %s\r\n", subject)
	message += fmt.Sprintf("\r\n%s\r\n", body)

	auth := smtp.PlainAuth("", config.GetConfig().Username, config.GetConfig().SmtpPassword, config.GetConfig().SmtpHost)
	err := smtp.SendMail(address, auth, config.GetConfig().Email, to, []byte(message))

	fmt.Println("email processed")
	return err

}
