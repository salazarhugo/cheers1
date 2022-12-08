package app

import (
	"bytes"
	"encoding/base64"
	"fmt"
	ticketpb "github.com/salazarhugo/cheers1/gen/go/cheers/ticket/v1"
	"github.com/salazarhugo/cheers1/services/email-service/internal/repository"
	"html/template"
	"log"
	"mime/multipart"
	"net/http"
	"net/smtp"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	host     = "smtp.gmail.com"
	username = "admin@maparty.fr"
	password = os.Getenv("EMAIL_PASSWORD")
	port     = "587"
)

type Message struct {
	To          []string
	CC          []string
	BCC         []string
	Subject     string
	Body        string
	Attachments map[string][]byte
}

type Sender struct {
	auth smtp.Auth
}

func New() *Sender {
	auth := smtp.PlainAuth("", username, password, host)

	return &Sender{auth}
}

func (s *Sender) Send(m *Message) error {
	return smtp.SendMail(fmt.Sprintf("%s:%s", host, port), s.auth, username, m.To, m.ToBytes())
}

func NewMessage(s, b string) *Message {
	return &Message{Subject: s, Body: b, Attachments: make(map[string][]byte)}
}

func (m *Message) AttachFileBytes(fileName string, b []byte) error {
	m.Attachments[fileName] = b
	return nil
}

func (m *Message) AttachFile(src string) error {
	b, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	_, fileName := filepath.Split(src)
	m.Attachments[fileName] = b
	return nil
}

func (m *Message) ToBytes() []byte {
	buf := bytes.NewBuffer(nil)
	withAttachments := len(m.Attachments) > 0
	buf.WriteString(fmt.Sprintf("Subject: %s\n", m.Subject))
	buf.WriteString(fmt.Sprintf("To: %s\n", strings.Join(m.To, ",")))
	if len(m.CC) > 0 {
		buf.WriteString(fmt.Sprintf("Cc: %s\n", strings.Join(m.CC, ",")))
	}

	if len(m.BCC) > 0 {
		buf.WriteString(fmt.Sprintf("Bcc: %s\n", strings.Join(m.BCC, ",")))
	}

	buf.WriteString("MIME-Version: 1.0\n")
	writer := multipart.NewWriter(buf)
	boundary := writer.Boundary()
	if withAttachments {
		buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n", boundary))
		buf.WriteString(fmt.Sprintf("--%s\n", boundary))
	} else {
		buf.WriteString("Content-Type: text/html; charset=utf-8\n")
	}

	buf.WriteString("Content-Type: text/html; charset=utf-8\n")
	buf.WriteString(m.Body)

	if withAttachments {
		for k, v := range m.Attachments {
			buf.WriteString(fmt.Sprintf("\n\n--%s\n", boundary))
			buf.WriteString(fmt.Sprintf("Content-Type: %s\n", http.DetectContentType(v)))
			buf.WriteString("Content-Transfer-Encoding: base64\n")
			buf.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=%s\n", k))

			b := make([]byte, base64.StdEncoding.EncodedLen(len(v)))
			base64.StdEncoding.Encode(b, v)
			buf.Write(b)
			buf.WriteString(fmt.Sprintf("\n--%s", boundary))
		}

		buf.WriteString("--")
	}

	return buf.Bytes()
}

func SendEmail(
	email string,
	firstName string,
	tickets []*ticketpb.Ticket,
	totalPrice int64,
) {
	amount := strconv.Itoa(int(totalPrice) / 100)
	body, err := ParseTemplate("internal/templates/payment-receipt.html", firstName, "", amount)
	if err != nil {
		log.Println(err)
		return
	}

	sender := New()
	m := NewMessage("Thank you for your payment", body)
	m.To = []string{email}
	m.CC = []string{"hugobrock74@gmail.com"}
	m.BCC = []string{}

	// Generate each ticket qrcode and attach it to email
	for i, ticket := range tickets {
		qrcode, err := repository.GenerateQRCodePng(ticket.Id)
		if err != nil {
			log.Println(err)
			return
		}
		m.AttachFileBytes(fmt.Sprintf("%s_%d", ticket.Name, i+1), qrcode)
	}

	err = sender.Send(m)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("Email Sent Successfully to %s!", email)
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
