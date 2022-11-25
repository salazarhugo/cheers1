package app

import (
	"bytes"
	"fmt"
	ticketpb "github.com/salazarhugo/cheers1/gen/go/cheers/ticket/v1"
	"html/template"
	"log"
	"net/mail"
	"net/smtp"
	"os"
	"strconv"
)

func SendEmail(
	email string,
	firstName string,
	tickets []*ticketpb.Ticket,
	totalPrice int64,
) {
	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Sender data.
	from := mail.Address{Name: "MaParty", Address: "admin@maparty.fr"}
	to := mail.Address{Address: email}

	pwd := os.Getenv("EMAIL_PASSWORD")
	amount := strconv.Itoa(int(totalPrice) / 100)

	body, err := ParseTemplate("internal/templates/payment-receipt.html", firstName, "", amount)
	if err != nil {
		log.Println(err)
		return
	}

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = "Thank you for your payment"
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=\"utf-8\""

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Authentication.
	auth := smtp.PlainAuth(
		"",
		from.Address,
		pwd,
		smtpHost,
	)

	// Sending email.
	err = smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth,
		from.Address,
		[]string{to.Address},
		[]byte(message),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Email Sent Successfully to %s!", to.Address)
}

func ParseTemplate(
	templateFileName string,
	firstName string,
	eventName string,
	price string,
) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		log.Println(err)
		return "", err
	}
	buf := new(bytes.Buffer)

	if err = t.Execute(buf, struct {
		Name      string
		EventName string
		Price     string
	}{
		Name:      firstName,
		EventName: eventName,
		Price:     price,
	}); err != nil {
		return "", err
	}

	return buf.String(), nil
}
