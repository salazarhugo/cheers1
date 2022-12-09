package app

import (
	payment "github.com/salazarhugo/cheers1/gen/go/cheers/payment/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"github.com/salazarhugo/cheers1/services/email-service/internal/repository"
	"log"
	"net/http"
)

func PaymentSub(w http.ResponseWriter, r *http.Request) {
	event := &payment.PaymentEvent{}
	err := pubsub.UnmarshalPubSubMessage(r, event)
	if err != nil {
		log.Fatal(err)
		return
	}

	order := event.Order
	totalPrice, err := repository.CalculateTotalPrice(order.Tickets)

	if err != nil {
		log.Println(err)
		http.Error(w, "failed to calculate amount", http.StatusInternalServerError)
		return
	}

	SendEmail(order.Email, order.FirstName, order.Tickets, totalPrice)
}
