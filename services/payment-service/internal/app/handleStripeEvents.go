package app

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/salazarhugo/cheers1/gen/go/cheers/payment/v1"
	utils2 "github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/payment-service/internal/repository"
	"github.com/salazarhugo/cheers1/services/payment-service/utils"
	"github.com/stripe/stripe-go/v72"
	"log"
	"net/http"
	"os"
)

func HandleStripeEvent(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	stripe.Key = os.Getenv("STRIPE_SK")

	event, err := repository.ValidateStripeRequest(cc)
	if err != nil {
		return err
	}

	switch event.Type {
	case "payment_intent.succeeded":
		var paymentIntent stripe.PaymentIntent
		err := parseEvent(event, &paymentIntent)
		if err != nil {
			return cc.NoContent(http.StatusInternalServerError)
		}
		log.Printf("Successful payment for %d.", paymentIntent.Amount)
		handlePaymentSuccess(paymentIntent)

	case "payment_method.attached":
		var paymentMethod stripe.PaymentMethod
		err := parseEvent(event, &paymentMethod)
		if err != nil {
			return cc.NoContent(http.StatusInternalServerError)
		}
	default:
		fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
	}

	return cc.NoContent(http.StatusOK)
}

func parseEvent(event *stripe.Event, v any) error {
	err := json.Unmarshal(event.Data.Raw, v)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
		return err
	}
	return nil
}

func handlePaymentSuccess(paymentIntent stripe.PaymentIntent) {
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
