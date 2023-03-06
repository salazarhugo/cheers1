package app

import (
	"github.com/google/uuid"
	activity2 "github.com/salazarhugo/cheers1/gen/go/cheers/activity/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"github.com/salazarhugo/cheers1/services/activity-service/internal/repository"
	"log"
	"net/http"
	"time"
)

func UserSub(w http.ResponseWriter, r *http.Request) {
	userEvent := &user.UserEvent{}

	err := pubsub.UnmarshalPubSubMessage(r, userEvent)
	if err != nil {
		log.Println(err)
		return
	}

	repo := repository.NewRepository()

	switch event := userEvent.Event.(type) {
	case *user.UserEvent_Follow:
		user := event.Follow.User
		activity := &activity2.Activity{
			Id:           uuid.New().String(),
			Type:         activity2.Activity_FOLLOW,
			Username:     user.Username,
			Text:         user.Username + " started following you",
			Picture:      user.Picture,
			UserId:       user.Id,
			Timestamp:    time.Now().Unix(),
			MediaPicture: "",
			MediaId:      "",
		}
		err = repo.CreateActivity(event.Follow.FollowedUser.Id, activity)
	case *user.UserEvent_Create:
	default:
	}

	if err != nil {
		log.Println(err)
		return
	}
}
