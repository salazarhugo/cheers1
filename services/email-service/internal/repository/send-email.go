package repository

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"os"
)

var (
	token = os.Getenv("SENDGRID_WEB_API_KEY")
)

func SendEmail(
	email string,
	name string,
	templateID string,
	data map[string]string,
) error {
	client := sendgrid.NewSendClient(token)

	from := mail.NewEmail("Cheers", "no-reply@maparty.app")
	to := mail.NewEmail(name, email)

	p := mail.NewPersonalization()
	p.AddTos(to)

	for key, value := range data {
		p.SetDynamicTemplateData(key, value)
	}

	message := mail.NewV3Mail()
	message.SetFrom(from)
	message.SetTemplateID(templateID)
	message.AddPersonalizations(p)

	_, err := client.Send(message)

	return err
}
