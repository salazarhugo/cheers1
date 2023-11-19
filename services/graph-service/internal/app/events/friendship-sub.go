package events

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/friendship/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"github.com/salazarhugo/cheers1/services/graph-service/internal/repository"
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

	repo := repository.NewUserRepository()

	switch event := event.Event.(type) {
	case *friendship.FriendshipEvent_CreatedFriendRequest:
		created := event.CreatedFriendRequest
		err = repo.CreateFriendRequest(created.From, created.To)
	case *friendship.FriendshipEvent_DeletedFriendRequest:
		deleted := event.DeletedFriendRequest
		err = repo.DeleteFriendRequest(deleted.From, deleted.To)
	case *friendship.FriendshipEvent_CreatedFriend:
		created := event.CreatedFriend
		err = repo.DeleteFriendRequest(created.From, created.To)
		err = repo.DeleteFriendRequest(created.To, created.From)
		err = repo.CreateFriend(created.From, created.To)
	case *friendship.FriendshipEvent_DeletedFriend:
		deleted := event.DeletedFriend
		err = repo.DeleteFriend(deleted.From, deleted.To)
	}

	if err != nil {
		http.Error(w, "failed to update order", http.StatusInternalServerError)
		return
	}

	return
}
