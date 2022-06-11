package main

import (
	"mail/service"
)

func main() {

	mailService := service.NewMailService()

	from := "cakmakforbusiness@gmail.com"
	to := "mertcakmak2@gmail.com"
	subject := "Subject: This is the subject of the mail\n"
	body := "This is the body of the mail"
	mailService.SendMail(from, to, subject, body)

}
