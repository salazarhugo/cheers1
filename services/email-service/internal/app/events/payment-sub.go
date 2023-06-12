package events

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

	//err = SendEmail(order.Email, order.FirstName, order.Tickets, totalPrice)
	err = SendEmail(
		order.Email,
		order.FirstName,
		"d-5b37c1f03f4d4551abbdc1f6b02d493f",
	)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(totalPrice)
}
