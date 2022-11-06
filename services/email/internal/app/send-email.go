package app

import (
	"fmt"
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
	// Message.
	message := []byte("Thank you for your payment of " + amount + " EUR")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
