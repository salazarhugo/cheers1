package app

import (
	"fmt"
	"github.com/google/uuid"
	activity2 "github.com/salazarhugo/cheers1/gen/go/cheers/activity/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/friendship/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"github.com/salazarhugo/cheers1/services/activity-service/internal/repository"
	"github.com/salazarhugo/cheers1/services/activity-service/internal/service"
	"log"
	"net/http"
	"time"
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
	case *friendship.FriendshipEvent_CreatedFriend:
		created := event.CreatedFriend
		friend1, err := service.GetUser(created.From)
		if err != nil {
			return
		}
		friend2, err := service.GetUser(created.To)
		if err != nil {
			return
		}
		activity := &activity2.Activity{
			Id:           uuid.New().String(),
			Type:         activity2.Activity_FRIEND_ADDED,
			Text:         fmt.Sprintf("%s added you as a friend.", friend2.Username),
			Username:     friend2.Username,
			Picture:      friend2.Picture,
			UserId:       friend2.Id,
			Timestamp:    time.Now().Unix(),
			MediaPicture: "",
			MediaId:      "",
		}
		err = repo.CreateActivity(friend1.Id, activity)
		activity2 := &activity2.Activity{
			Id:           uuid.New().String(),
			Type:         activity2.Activity_FRIEND_ADDED,
			Text:         fmt.Sprintf("%s added you as a friend.", friend1.Username),
			Username:     friend1.Username,
			Picture:      friend1.Picture,
			UserId:       friend1.Id,
			Timestamp:    time.Now().Unix(),
			MediaPicture: "",
			MediaId:      "",
		}
		err = repo.CreateActivity(friend2.Id, activity2)
	case *friendship.FriendshipEvent_CreatedFriendRequest:
	case *friendship.FriendshipEvent_DeletedFriendRequest:
	case *friendship.FriendshipEvent_DeletedFriend:
	}

	if err != nil {
		http.Error(w, "failed to update order", http.StatusInternalServerError)
		return
	}

	return
}
