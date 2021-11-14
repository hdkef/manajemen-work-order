package services

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "Email dari Server <hadekha.edu@gmail.com>"

var CONFIG_AUTH_EMAIL string
var CONFIG_AUTH_PASS string

func init() {
	_ = godotenv.Load()
	CONFIG_AUTH_EMAIL = os.Getenv("CONFIGAUTHEMAIL")
	CONFIG_AUTH_PASS = os.Getenv("CONFIGAUTHPASS")
}

func SendEmail(toWhom string, subject string, message string) error {
	to := []string{toWhom}

	body := "Subject: " + subject + "\n\n" + message
	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASS, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	return smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, to, []byte(body))
}
