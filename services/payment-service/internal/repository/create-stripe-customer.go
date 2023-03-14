package repository

import (
	"fmt"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/customer"
)

func (r paymentRepository) CreateStripeCustomer(
	userID string,
	email string,
) (string, error) {
	// Create a new customer with the given email address and metadata.
	params := &stripe.CustomerParams{
		Email: stripe.String(email),
		Params: stripe.Params{
			Metadata: map[string]string{
				"userID": userID,
			},
		},
	}

	cust, err := customer.New(params)
	if err != nil {
		return "", fmt.Errorf("failed to create stripe customer: %v", err)
	}

	// Return the customer ID.
	return cust.ID, nil
}
