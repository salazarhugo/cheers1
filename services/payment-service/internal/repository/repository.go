package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/salazarhugo/cheers1/libs/utils"
)

type PaymentRepository interface {
	CreatePayment(userID string) (string, error)
	CreateStripeCustomer(userID string, email string) (string, error)
}

type paymentRepository struct {
	client *firestore.Client
}

func NewPaymentRepository() PaymentRepository {
	app := utils.InitializeAppDefault()
	client, _ := app.Firestore(context.Background())
	return &paymentRepository{
		client: client,
	}
}
