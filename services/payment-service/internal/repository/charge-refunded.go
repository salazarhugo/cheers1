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

	customerId := ""
	if paymentIntent.Customer != nil {
		customerId = paymentIntent.Customer.ID
	}

	err := pubsub.PublishProtoWithBinaryEncoding("payment-topic", &payment.PaymentEvent{
		PaymentIntentId: paymentIntent.ID,
		CustomerId:      customerId,
		Type:            payment.PaymentEvent_REFUND,
	})
	if err != nil {
		log.Println(err)
	}
}
