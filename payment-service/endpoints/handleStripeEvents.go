package endpoints

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/webhook"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"salazar/cheers/payment/utils"
)

func HandleStripeEvent(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()
	stripe.Key = os.Getenv("STRIPE_SK")

	const MaxBodyBytes = int64(65536)
	cc.Request().Body = http.MaxBytesReader(cc.Response().Writer, cc.Request().Body, MaxBodyBytes)
	payload, err := ioutil.ReadAll(cc.Request().Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
		return cc.NoContent(http.StatusServiceUnavailable)
	}

	event := stripe.Event{}

	if err := json.Unmarshal(payload, &event); err != nil {
		fmt.Fprintf(os.Stderr, "⚠️  Webhook error while parsing basic request. %v\n", err.Error())
		return cc.NoContent(http.StatusBadRequest)
	}

	endpointSecret := os.Getenv("STRIPE_WH")
	signatureHeader := cc.Request().Header.Get("Stripe-Signature")
	event, err = webhook.ConstructEvent(payload, signatureHeader, endpointSecret)

	if err != nil {
		fmt.Fprintf(os.Stderr, "⚠️  Webhook signature verification failed. %v\n", err)
		return cc.NoContent(http.StatusBadRequest)
	}

	switch event.Type {
	case "payment_intent.succeeded":
		var paymentIntent stripe.PaymentIntent
		err := json.Unmarshal(event.Data.Raw, &paymentIntent)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			return cc.NoContent(http.StatusBadRequest)
		}
		log.Printf("Successful payment for %d.", paymentIntent.Amount)
		handleSuccess(paymentIntent)

	case "payment_method.attached":
		var paymentMethod stripe.PaymentMethod
		err := json.Unmarshal(event.Data.Raw, &paymentMethod)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			return cc.NoContent(http.StatusBadRequest)
		}
	default:
		fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
	}

	return cc.NoContent(http.StatusOK)
}

func handleSuccess(paymentIntent stripe.PaymentIntent) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "cheers-a275e")
	if err != nil {
		return
	}

	customerId := paymentIntent.Customer.ID
	snapshot := client.Collection("stripe_customers").Where("customer_id", "==", customerId).Snapshots(ctx)

	docs, err := snapshot.Next()
	customerDoc, err := docs.Documents.Next()
	customerRef := customerDoc.Ref

	_, _, err = customerRef.Collection("payments").Add(ctx, map[string]interface{}{
		"id":         paymentIntent.ID,
		"amount":     paymentIntent.Amount,
		"created":    paymentIntent.Created,
		"customerId": paymentIntent.Customer.ID,
		"status":     paymentIntent.Status,
	})

	docsnap, err := customerRef.Collection("paymentIntents").Doc(paymentIntent.ID).Get(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	docMap := docsnap.Data()
	partyId := docMap["partyId"]
	_, err = customerRef.Collection("private").Doc(customerId).Update(ctx, []firestore.Update{
		{Path: "population", Value: firestore.Increment(paymentIntent.Amount)},
	})

	log.Println(partyId)
	if err != nil {
		log.Println(err)
		return
	}
}
