package app

import (
	payment "github.com/salazarhugo/cheers1/gen/go/cheers/payment/v1"
	"github.com/salazarhugo/cheers1/services/account-service/internal/repository"
	"log"
	"net/http"
)

func PaymentSub(w http.ResponseWriter, r *http.Request) {
	event := &payment.PaymentEvent{}
	err := ParsePubSubEvent(w, r, event)
	if err != nil {
		http.Error(w, "failed to parse pub/sub event", http.StatusInternalServerError)
	}

	switch event.Type {
	case payment.PaymentEvent_REFUND:
	case payment.PaymentEvent_PAYMENT_SUCCESS:
	}

	log.Println(event)

	order, err := repository.GetOrder(event.PaymentIntentId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	repository := repository.NewAccountRepository()
	err = repository.IncrementBalance(order.PartyHostId, order.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}
