package events

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/auth/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"github.com/salazarhugo/cheers1/services/payment-service/internal/repository"
	"log"
	"net/http"
)

func AuthSub(w http.ResponseWriter, r *http.Request) {
	event := &auth.AuthEvent{}
	err := pubsub.UnmarshalPubSubMessage(r, event)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println(event)

	repo := repository.NewPaymentRepository()

	switch event := event.Event.(type) {
	case *auth.AuthEvent_Created:
		_, err = repo.CreateStripeCustomer(event.Created.Uid, event.Created.Email)
	}

	if err != nil {
		log.Println(err)
		http.Error(w, "failed to handle event: "+err.Error(), http.StatusInternalServerError)
		return
	}

	return
}
