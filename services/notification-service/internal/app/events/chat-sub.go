package events

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
		room := event.Create.Room
		for _, member := range members {
			// Don't send notification to the sender
			if member.Id == sender.Id {
				continue
			}
			go Async(
				repo,
				sender,
				room,
				member,
				event.Create.Message.RoomId,
			)
		}
	}

	return
}

func Async(
	repo repository.Repository,
	sender *user.UserItem,
	room *chat.Room,
	member *user.UserItem,
	roomId string,
) {
	tokens, err := repo.GetUserTokens(member.Id)
	if err != nil {
		log.Println(err)
		return
	}

	notification := make(map[string]string, 0)

	switch room.Type {
	case chat.RoomType_DIRECT:
		notification = notifications.NewChatMessageNotification(
			sender.Username,
			sender.Picture,
			roomId,
		)
	case chat.RoomType_GROUP:
		notification = notifications.NewGroupMessageNotification(
			room.Name,
			room.Picture,
			roomId,
		)
	}

	err = repo.SendNotification(
		map[string][]string{member.Id: tokens},
		notification,
	)

	if err != nil {
		log.Println(err)
		return
	}
}
