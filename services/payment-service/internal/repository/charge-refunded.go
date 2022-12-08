package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/payment/v1"
	utils2 "github.com/salazarhugo/cheers1/libs/utils"
	"github.com/stripe/stripe-go/v72"
	"log"
)

func HandleChargeRefunded(charge stripe.Charge) {
	log.Printf("Successful refund of %d.", charge.AmountRefunded)
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "cheers-a275e")
	if err != nil {
		return
	}

	paymentIntent := charge.PaymentIntent

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
			Type:            payment.PaymentEvent_REFUND,
		})
	}()

	if paymentIntent.Customer != nil {
		snapshot := client.Collection("stripe_customers").Where("customer_id", "==", customerId).Snapshots(ctx)
		docs, err := snapshot.Next()
		stripeCustomerDoc, err := docs.Documents.Next()
		stripeCustomerRef := stripeCustomerDoc.Ref

		if err != nil {
			log.Println(err)
			return
		}
	}
}
