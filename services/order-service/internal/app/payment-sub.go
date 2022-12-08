package app

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/order/v1"
	payment "github.com/salazarhugo/cheers1/gen/go/cheers/payment/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"github.com/salazarhugo/cheers1/services/order-service/internal/repository"
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

	repo := repository.NewOrderRepository()

	switch event.Type {
	case payment.PaymentEvent_PAYMENT_SUCCESS:
		err = repo.UpdateOrder(&pb.Order{
			Status: "Succeeded",
			Id:     event.PaymentIntentId,
		})
	case payment.PaymentEvent_REFUND:
		err = repo.UpdateOrder(&pb.Order{
			Status: "Refunded",
			Id:     event.PaymentIntentId,
		})
	}

	if err != nil {
		http.Error(w, "failed to update order", http.StatusInternalServerError)
		return
	}

	return
}
