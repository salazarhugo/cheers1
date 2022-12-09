package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/iterator"
)

func (o ticketRepository) DeleteTicket(
	userID string,
	orderID string,
) error {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, "cheers-a275e")
	if err != nil {
		return err
	}

	docs := client.Collection("users").Doc(userID).Collection("tickets").Where("paymentIntentId", "==", orderID).Documents(ctx)
	for {
		doc, err := docs.Next()
		if err == iterator.Done {
			break
		}
		doc.Ref.Delete(ctx)
	}

	docs = client.Collection("tickets").Where("paymentIntentId", "==", orderID).Documents(ctx)
	for {
		doc, err := docs.Next()
		if err == iterator.Done {
			break
		}
		doc.Ref.Delete(ctx)
	}

	return nil
}
