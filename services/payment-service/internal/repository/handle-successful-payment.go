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

	// Retrieve the order of that payment
	order, err := GetOrder(paymentIntent.ID)
	if err != nil {
		log.Println(err)
		return
	}

	customerId := ""
	if paymentIntent.Customer != nil {
		customerId = paymentIntent.Customer.ID
	}

	go func() {
		err = utils2.PublishProtoMessages("payment-topic", &payment.PaymentEvent{
			PaymentIntentId: paymentIntent.ID,
			CustomerId:      customerId,
			Type:            payment.PaymentEvent_PAYMENT_SUCCESS,
			FirstName:       order.FirstName,
			LastName:        order.LastName,
			Email:           order.Email,
		})
	}()

	if paymentIntent.Customer != nil {
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

		for _, ticket := range order.Tickets {
			doc := client.Collection("tickets").NewDoc()
			ticket.Id = doc.ID
			ticket.PaymentIntentId = paymentIntent.ID
			ticket.UserId = stripeCustomerRef.ID
			ticket.Validated = false
			m, _ := utils2.ProtoToMap(ticket)
			_, err = doc.Set(ctx, m)
			_, err = client.Collection("users").Doc(stripeCustomerRef.ID).Collection("tickets").Doc(doc.ID).Set(ctx, m)
		}

		if err != nil {
			log.Println(err)
			return
		}
	}
}
