package events

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/friendship/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"github.com/salazarhugo/cheers1/services/notification-service/internal/notifications"
	"github.com/salazarhugo/cheers1/services/notification-service/internal/repository"
	"github.com/salazarhugo/cheers1/services/notification-service/internal/service"
	"log"
	"net/http"
)

func FriendShipSub(w http.ResponseWriter, r *http.Request) {
	event := &friendship.FriendshipEvent{}
	err := pubsub.UnmarshalPubSubMessage(r, event)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println(event)

	repo := repository.NewRepository()

	switch event := event.Event.(type) {
	case *friendship.FriendshipEvent_CreatedFriendRequest:
		created := event.CreatedFriendRequest
		user, err := service.GetUser(created.From)
		if err != nil {
			return
		}
		tokens, err := repo.GetUserTokens(created.To)
		if err != nil {
			return
		}

		data := notifications.FriendRequestNotification(user.Username, user.Picture)
		err = repo.SendNotification(map[string][]string{created.To: tokens}, data)
	case *friendship.FriendshipEvent_DeletedFriendRequest:
	case *friendship.FriendshipEvent_CreatedFriend:
		created := event.CreatedFriend
		user, err := service.GetUser(created.From)
		if err != nil {
			return
		}
		tokens, err := repo.GetUserTokens(created.To)
		if err != nil {
			return
		}
		data := notifications.AcceptedFriendRequestNotification(user.Username, user.Picture)
		err = repo.SendNotification(map[string][]string{created.To: tokens}, data)
	case *friendship.FriendshipEvent_DeletedFriend:
	}

	if err != nil {
		http.Error(w, "failed to update order", http.StatusInternalServerError)
		return
	}

	return
}
