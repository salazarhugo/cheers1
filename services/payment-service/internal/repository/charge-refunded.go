package repository

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/payment/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"github.com/stripe/stripe-go/v72"
	"log"
)

func HandleChargeRefunded(charge stripe.Charge) {
	log.Printf("Successful refund of %d.", charge.AmountRefunded)

	paymentIntent := charge.PaymentIntent

	// Retrieve the order of that payment
	order, err := GetOrder(paymentIntent.ID)
	if err != nil {
		log.Println(err)
		return
	}

	err = pubsub.PublishProtoWithBinaryEncoding("payment-topic", &payment.PaymentEvent{
		PaymentIntentId: paymentIntent.ID,
		Type:            payment.PaymentEvent_REFUND,
		Order:           order,
	})
	if err != nil {
		log.Println(err)
	}
}
