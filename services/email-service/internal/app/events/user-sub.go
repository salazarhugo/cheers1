package events

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"github.com/salazarhugo/cheers1/services/email-service/internal/emails"
	"log"
	"net/http"
)

func UserSub(w http.ResponseWriter, r *http.Request) {
	event := &user.UserEvent{}
	err := pubsub.UnmarshalPubSubMessage(r, event)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println(event)

	switch event := event.Event.(type) {
	case *user.UserEvent_Create:
		log.Println(event.Create)
		user := event.Create.User
		err := emails.SendWelcomeEmail(user.Email, user.Name)
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
