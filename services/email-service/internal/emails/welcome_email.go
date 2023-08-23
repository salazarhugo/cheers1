package emails

import "github.com/salazarhugo/cheers1/services/email-service/internal/repository"

func SendWelcomeEmail(
	email string,
	name string,
) error {
	return repository.SendEmail(
		email,
		name,
		"d-636b81638b4e4f068c3f181a450c3b13",
		map[string]string{
			"first_name": name,
		},
	)
}
