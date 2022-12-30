package app

import (
	chat "github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
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

	log.Println(event)

	repo := repository.NewRepository()

	switch event := event.Event.(type) {
	case *chat.ChatEvent_Create:
		members := event.Create.Members
		sender := event.Create.Sender
		for _, member := range members {
			// Don't send notification to the sender
			if member.Id == sender.Id {
				continue
			}
			go Async(repo, sender, member, event.Create.Message.RoomId)
		}
	}

	return
}

func Async(
	repo repository.Repository,
	sender *user.UserItem,
	member *user.UserItem,
	roomId string,
) {
	tokens, err := repo.GetUserTokens(member.Id)
	if err != nil {
		log.Println(err)
		return
	}

	data := notifications.NewChatMessageNotification(
		sender.Username,
		sender.Picture,
		roomId,
	)

	err = repo.SendNotification(
		map[string][]string{member.Id: tokens},
		data,
	)

	if err != nil {
		log.Println(err)
		return
	}
}
