package app

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"github.com/salazarhugo/cheers1/services/auth-service/internal/repository"
	"log"
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
		user := event.Create.User
		if user.Email != "hugobrock74@gmail.com" && user.Email != "admin@maparty.app" {
			return
		}
		err = repo.CreateAdmin(user.Id)
		err = repo.CreateModerator(user.Id)
		err = repo.VerifyUser(user.Id)
	}

	if err != nil {
		log.Println(err)
		http.Error(w, "failed to verify user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	return
}
