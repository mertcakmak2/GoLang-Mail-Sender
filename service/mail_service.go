package service

import (
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

type MailService struct{}

func NewMailService() MailService {
	return MailService{}
}

func (m MailService) SendMail(from, toEmailAddress, subject, body string) {

	password := getEnv("PASSWORD")
	to := []string{toEmailAddress}
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	message := []byte(subject + body)
	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		panic(err)
	}

}

func getEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dotenv := os.Getenv(key)
	return dotenv
}
