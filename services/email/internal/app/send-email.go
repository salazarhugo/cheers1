package app

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"strconv"
)

func SendEmail(
	email string,
	tickets []map[string]interface{},
	totalPrice int64,
) {
	log.Println(email)
	log.Println(tickets)
	log.Println(totalPrice)

	// Sender data.
	from := "admin@maparty.fr"
	pwd := os.Getenv("EMAIL_PASSWORD")
	password := pwd

	// Receiver email address.
	to := []string{
		email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	amount := strconv.Itoa(int(totalPrice) / 100)

	body, err := ParseTemplate("internal/templates/payment-receipt.html")
	if err != nil {
		log.Println(err)
		return
	}

	subject := "Subject: Thank you for your payment of " + amount + " EUR\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	// Message.
	message := []byte(subject + mime + body)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}

func ParseTemplate(templateFileName string) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		log.Println(err)
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, nil); err != nil {
		return "", err
	}

	return buf.String(), nil
}
