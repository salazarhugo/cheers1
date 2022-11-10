package repository

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"github.com/salazarhugo/cheers1/libs/utils"
)

func GetUserId(customerId string) (string, error) {
	ctx := context.Background()
	app := utils.InitializeAppDefault()
	client, err := app.Firestore(ctx)
	if err != nil {
		return "", err
	}

	snapshot := client.Collection("stripe_customers").Where("customer_id", "==", customerId).Documents(ctx)
	doc, err := snapshot.Next()
	if err != nil {
		return "", err
	}

	return doc.Ref.ID, nil
}

func GetAuthUser(userID string) (*auth.UserRecord, error) {
	ctx := context.Background()
	app := utils.InitializeAppDefault()

	auth, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}
	user, err := auth.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetOrder(paymentIntentId string, customerId string) ([]map[string]interface{}, error) {
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

	tickets := orderMap["tickets"].([]interface{})
	var res = make([]map[string]interface{}, 0)

	for _, ticket := range tickets {
		res = append(res, ticket.(map[string]interface{}))
	}

	return res, nil
}

func CalculateTotalPrice(tickets []map[string]interface{}) int64 {
	var total int64

	for _, ticket := range tickets {
		price, ok := ticket["price"].(int64)
		if ok {
			total += price
		}
	}

	return total
}
