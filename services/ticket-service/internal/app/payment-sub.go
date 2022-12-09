package app

import (
	payment "github.com/salazarhugo/cheers1/gen/go/cheers/payment/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"github.com/salazarhugo/cheers1/services/ticket-service/internal/repository"
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
	log.Println(event)

	repo := repository.NewTicketRepository()

	switch event.Type {
	case payment.PaymentEvent_PAYMENT_SUCCESS:
		repo.CreateTicket(event.CustomerId, event.)
	case payment.PaymentEvent_REFUND:
	}

	if err != nil {
		http.Error(w, "failed to update ticket", http.StatusInternalServerError)
		return
	}

	return
}
