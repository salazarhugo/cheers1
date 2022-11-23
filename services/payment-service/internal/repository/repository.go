package repository

import (
	"cloud.google.com/go/firestore"
)

type PaymentRepository interface {
	CreatePayment(userID string) (string, error)
}

type paymentRepository struct {
	client *firestore.Client
}

func NewPaymentRepository(client *firestore.Client) PaymentRepository {
	return &paymentRepository{
		client: client,
	}
}
