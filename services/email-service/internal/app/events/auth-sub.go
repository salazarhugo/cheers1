package events

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/auth/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
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

	switch event := event.Event.(type) {
	case *auth.AuthEvent_Created:
		log.Println(event.Created)
		//_, err = repository.GetOrder(event.Created.Uid, event.Created.Email)
	}

	if err != nil {
		log.Println(err)
		http.Error(w, "failed to handle event: "+err.Error(), http.StatusInternalServerError)
		return
	}

	return
}
