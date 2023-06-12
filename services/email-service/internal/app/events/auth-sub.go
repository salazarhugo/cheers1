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
		user := event.Created
		err := SendEmail(user.Email, "Test", "d-636b81638b4e4f068c3f181a450c3b13")
		if err != nil {
			log.Println(err)
			return
		}
	}

	if err != nil {
		log.Println(err)
		http.Error(w, "failed to handle event: "+err.Error(), http.StatusInternalServerError)
		return
	}

	return
}
