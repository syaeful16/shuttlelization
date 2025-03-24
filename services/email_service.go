package services

import (
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

type EmailService struct {
	SMTPHost string
	SMTPPort string
	Username string
	Password string
	From     string
}

func NewEmailService() *EmailService {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Gagal membaca file .env")
	}

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")
	from := os.Getenv("SMPT_FROM")

	return &EmailService{
		SMTPHost: smtpHost,
		SMTPPort: smtpPort,
		Username: username,
		Password: password,
		From:     from,
	}
}

func (e *EmailService) SendEmail(to, subject, body string) error {
	auth := smtp.PlainAuth("", e.Username, e.Password, e.SMTPHost)

	// Format pesan email
	message := []byte("Subject: " + subject + "\r\n" +
		"From: " + e.From + "\r\n" +
		"To: " + to + "\r\n" +
		"\r\n" +
		body + "\r\n")

	// Kirim email
	err := smtp.SendMail(e.SMTPHost+":"+e.SMTPPort, auth, e.From, []string{to}, message)
	if err != nil {
		return err
	}

	return nil
}
