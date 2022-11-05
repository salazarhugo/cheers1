package repository

import (
	"context"
	"github.com/salazarhugo/cheers1/libs/utils"
	"log"
)

func GetEmailAndOrder(paymentIntentId string, customerId string) (*string, error) {
	ctx := context.Background()
	app := utils.InitializeAppDefault()
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	doc, err := client.Collection("orders").Doc(paymentIntentId).Get(ctx)
	if err != nil {
		return nil, err
	}
	orderMap := doc.Data()
	tickets := orderMap["tickets"]
	log.Println(tickets)

	snapshot := client.Collection("stripe_customers").Where("customer_id", "==", customerId).Documents(ctx)
	doc, err = snapshot.Next()
	if err != nil {
		return nil, err
	}

	userID := doc.Ref.ID

	auth, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}
	user, err := auth.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &user.Email, nil
}
