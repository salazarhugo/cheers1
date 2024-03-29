package app

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"github.com/salazarhugo/cheers1/services/comment-service/internal/repository"
	"net/http"
)

func UserSub(w http.ResponseWriter, r *http.Request) {
	event := &user.UserEvent{}
	err := pubsub.UnmarshalPubSubMessage(r, event)
	if err != nil {
		return
	}

	repo := repository.NewRepository()

	switch event := event.Event.(type) {
	case *user.UserEvent_Create:
		err = repo.UpdateUser(event.Create.User)
	case *user.UserEvent_Update:
		err = repo.UpdateUser(event.Update.User)
	}

	if err != nil {
		http.Error(w, "failed to update order", http.StatusInternalServerError)
		return
	}

	return
}
