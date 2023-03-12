package app

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/auth/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"github.com/salazarhugo/cheers1/services/user-service/internal/domain"
	"log"
	"net/http"
)

func AuthSub(w http.ResponseWriter, r *http.Request) {
	event := &auth.AuthEvent{}
	err := pubsub.UnmarshalPubSubMessage(r, event)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println(event)

	userService := domain.NewUserService()

	switch event := event.Event.(type) {
	case *auth.AuthEvent_Deleted:
		userID := event.Deleted.Uid
		err = userService.DeleteUser(userID)
	}

	if err != nil {
		log.Println(err)
		http.Error(w, "failed to delete user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	return
}
