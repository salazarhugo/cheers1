package emails

import (
	"github.com/salazarhugo/cheers1/services/email-service/internal/repository"
)

func SendPartyReceiptEmail(
	email string,
	name string,
	partyName string,
) error {
	return repository.SendEmail(
		email,
		name,
		"d-5b37c1f03f4d4551abbdc1f6b02d493f",
		map[string]string{
			"first_name": name,
			"party_name": partyName,
		},
	)
}
