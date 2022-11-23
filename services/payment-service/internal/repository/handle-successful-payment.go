package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/payment/v1"
	utils2 "github.com/salazarhugo/cheers1/libs/utils"
	"github.com/stripe/stripe-go/v72"
	"log"
)

func HandlePaymentSuccess(paymentIntent stripe.PaymentIntent) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "cheers-a275e")
	if err != nil {
		return
	}

	customerId := paymentIntent.Customer.ID
	snapshot := client.Collection("stripe_customers").Where("customer_id", "==", customerId).Snapshots(ctx)

	docs, err := snapshot.Next()
	stripeCustomerDoc, err := docs.Documents.Next()
	stripeCustomerRef := stripeCustomerDoc.Ref

	// Store payment details into firestore
	_, _, err = stripeCustomerRef.Collection("payments").Add(ctx, map[string]interface{}{
		"id":         paymentIntent.ID,
		"amount":     paymentIntent.Amount,
		"created":    paymentIntent.Created,
		"customerId": paymentIntent.Customer.ID,
		"status":     paymentIntent.Status,
	})

	// Retrieve the order of that payment
	orderDoc, err := client.Collection("orders").Doc(paymentIntent.ID).Get(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	orderDocMap := orderDoc.Data()
	//partyId := orderDocMap["partyId"].(string)
	tickets := orderDocMap["tickets"].([]interface{})

	for _, ticket := range tickets {
		ticketMap := ticket.(map[string]interface{})
		doc := client.Collection("tickets").NewDoc()

		ticketData := map[string]interface{}{
			"id":              doc.ID,
			"name":            ticketMap["name"],
			"description":     ticketMap["description"],
			"price":           ticketMap["price"],
			"partyId":         ticketMap["partyId"],
			"paymentIntentId": paymentIntent.ID,
			"userId":          stripeCustomerRef.ID,
			"validated":       false,
		}

		_, err = doc.Set(ctx, ticketData)
		_, err = client.Collection("users").Doc(stripeCustomerRef.ID).Collection("tickets").Doc(doc.ID).Set(ctx, ticketData)
	}

	_, err = stripeCustomerRef.Collection("private").Doc(customerId).Update(ctx, []firestore.Update{
		{Path: "population", Value: firestore.Increment(paymentIntent.Amount)},
	})

	err = utils2.PublishProtoMessages("payment-topic", &payment.PaymentEvent{
		PaymentIntentId: paymentIntent.ID,
		CustomerId:      customerId,
		Type:            payment.PaymentEvent_PAYMENT_SUCCESS,
	})

	if err != nil {
		log.Println(err)
		return
	}
}
