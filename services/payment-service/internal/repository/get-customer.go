package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StripeCustomer struct {
	Id string `firestore:"customer_id"`
}

func GetCustomer(userID string) (*StripeCustomer, error) {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, "cheers-a275e")
	if err != nil {
		return nil, err
	}

	userDoc, err := client.Collection("stripe_customers").Doc(userID).Get(ctx)
	if err != nil {
		return nil, err
	}

	var customer StripeCustomer
	if err := userDoc.DataTo(&customer); err != nil {
		return nil, status.Error(codes.NotFound, "customer not found")
	}

	return &customer, nil
}
