package main

import (
	"fmt"
	"net/smtp"
	"os"
)

func sendEmailHTML(subject string, html string, to []string) {
	auth := smtp.PlainAuth(
		"",
		"arunarivanandam@gmail.com",
		os.Getenv("APP_PASSWORD"),
		"smtp.gmail.com",
	)
	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := "Subject:" + subject + "\n" + headers + html
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"arunarivanandam@gmail.com",
		to,
		[]byte(msg),
	)
	if err != nil {
		fmt.Println(err)
	}
}