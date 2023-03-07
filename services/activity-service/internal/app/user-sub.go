package app

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"log"
	"net/http"
)

func UserSub(w http.ResponseWriter, r *http.Request) {
	userEvent := &user.UserEvent{}

	err := pubsub.UnmarshalPubSubMessage(r, userEvent)
	if err != nil {
		log.Println(err)
		return
	}
}
