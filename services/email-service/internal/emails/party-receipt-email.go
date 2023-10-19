package emails

import (
	"github.com/salazarhugo/cheers1/services/email-service/internal/repository"
)

type EmailTicket struct {
	name     string
	quantity uint32
	price    int64
}

func New(name string, quantity uint32, price int64) *EmailTicket {
	return &EmailTicket{name, quantity, price}
}

func SendPartyReceiptEmail(
	email string,
	name string,
	partyName string,
	tickets []*EmailTicket,
) error {
	return repository.SendEmail(
		email,
		name,
		"d-5b37c1f03f4d4551abbdc1f6b02d493f",
		map[string]interface{}{
			"first_name": name,
			"party_name": partyName,
			"tickets":    tickets,
		},
	)
}
