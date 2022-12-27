package app

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"github.com/salazarhugo/cheers1/services/notification-service/internal/repository"
	"log"
	"net/http"
)

func UserTopicSub(w http.ResponseWriter, r *http.Request) {
	userEvent := &user.UserEvent{}

	err := pubsub.UnmarshalPubSubMessage(r, userEvent)
	if err != nil {
		log.Println(err)
		return
	}

	repo := repository.NewRepository()

	switch event := userEvent.Event.(type) {
	case *user.UserEvent_Follow:
		err = repo.FollowUserNotification(event.Follow.User, event.Follow.FollowedUser)
	case *user.UserEvent_Create:
	default:
	}

	if err != nil {
		log.Println(err)
		return
	}
}
