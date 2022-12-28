package app

import (
	chat "github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"github.com/salazarhugo/cheers1/services/activity-service/internal/repository"
	"log"
	"net/http"
)

func ChatSub(w http.ResponseWriter, r *http.Request) {
	event := &chat.ChatEvent{}
	err := pubsub.UnmarshalPubSubMessage(r, event)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(event)

	repo := repository.NewRepository()
	log.Println(repo)

	switch event := event.Event.(type) {
	case *chat.ChatEvent_Create:
		members := event.Create.Members
		sender := event.Create.Sender
		for _, member := range members {
			// Don't send notification to the sender
			if member.Id == sender.Id {
				continue
			}
		}
	}

	return
}
