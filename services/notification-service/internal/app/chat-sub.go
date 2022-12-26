package app

import (
	chat "github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"github.com/salazarhugo/cheers1/services/notification-service/internal/notifications"
	"github.com/salazarhugo/cheers1/services/notification-service/internal/repository"
	"log"
	"net/http"
)

func ChatTopic(w http.ResponseWriter, r *http.Request) {
	event := &chat.ChatEvent{}
	err := pubsub.UnmarshalPubSubMessage(r, event)
	if err != nil {
		log.Println(err)
		return
	}

	repo := repository.NewRepository()

	switch event := event.Event().(type) {
	case *chat.ChatEvent_Create:
		members := event.Create.Members
		for _, member := range members {
			tokens, err := repo.GetUserTokens(member.Id)
			if err != nil {
				return
			}

			data := notifications.NewChatMessageNotification(member.Username, member.Picture)
			go repo.SendNotification(map[string][]string{
				member.Id: tokens,
			}, data)
		}
	}

	return
}
