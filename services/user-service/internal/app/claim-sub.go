package app

import (
	claim "github.com/salazarhugo/cheers1/gen/go/cheers/claim/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"github.com/salazarhugo/cheers1/services/user-service/internal/repository"
	"log"
	"net/http"
)

func ClaimSub(w http.ResponseWriter, r *http.Request) {
	event := &claim.ClaimEvent{}
	err := pubsub.UnmarshalPubSubMessage(r, event)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println(event)

	repo := repository.NewUserRepository()

	switch event := event.Event.(type) {
	case *claim.ClaimEvent_Created:
		if event.Created.Claim == "verified" {
			err = repo.VerifyUser(event.Created.UserId)
		}
	case *claim.ClaimEvent_Deleted:
		if event.Deleted.Claim == "verified" {
			err = repo.UnVerifyUser(event.Deleted.UserId)
		}
	}

	if err != nil {
		http.Error(w, "failed to update order", http.StatusInternalServerError)
		return
	}

	return
}
