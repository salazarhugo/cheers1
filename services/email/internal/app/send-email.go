package app

import (
	"bytes"
	"firebase.google.com/go/v4/auth"
	"fmt"
	"html/template"
	"log"
	"net/mail"
	"net/smtp"
	"os"
	"strconv"
)

func SendEmail(
	user *auth.UserRecord,
	tickets []map[string]interface{},
	totalPrice int64,
) {
	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Sender data.
	from := mail.Address{Name: "MaParty", Address: "admin@maparty.fr"}
	to := mail.Address{Address: user.Email}

	pwd := os.Getenv("EMAIL_PASSWORD")
	amount := strconv.Itoa(int(totalPrice) / 100)

	body, err := ParseTemplate("internal/templates/payment-receipt.html", user, amount)
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

	log.Println(message)

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
	fmt.Println("Email Sent Successfully!")
}

func ParseTemplate(
	templateFileName string,
	user *auth.UserRecord,
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
		Name:      user.DisplayName,
		EventName: "Summer Breeze",
		Price:     price,
	}); err != nil {
		return "", err
	}

	return buf.String(), nil
}
