package app

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/repository"
	"net/http"
)

func UserSub(w http.ResponseWriter, r *http.Request) {
	event := &user.UserEvent{}
	err := pubsub.UnmarshalPubSubMessage(r, event)
	if err != nil {
		return
	}

	repo := repository.NewChatRepository()

	switch event := event.Event.(type) {
	case *user.UserEvent_Delete:
		err = repo.DeleteRooms(event.Delete.UserId)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}
