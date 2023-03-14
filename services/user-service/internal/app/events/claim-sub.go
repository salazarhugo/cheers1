package events

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
		claimName := event.Created.Claim
		if claimName == "verified" {
			err = repo.VerifyUser(event.Created.UserId)
		}
		if claimName == "business" {
			err = repo.UpdateBusinessAccount(event.Created.UserId, true)
		}
		if claimName == "moderator" {
			err = repo.UpdateModerator(event.Created.UserId, true)
		}
		if claimName == "admin" {
			err = repo.UpdateAdmin(event.Created.UserId, true)
		}
	case *claim.ClaimEvent_Deleted:
		claimName := event.Deleted.Claim
		if claimName == "verified" {
			err = repo.UnVerifyUser(event.Deleted.UserId)
		}
		if claimName == "business" {
			err = repo.UpdateBusinessAccount(event.Deleted.UserId, false)
		}
		if claimName == "moderator" {
			err = repo.UpdateModerator(event.Deleted.UserId, false)
		}
		if claimName == "admin" {
			err = repo.UpdateAdmin(event.Deleted.UserId, false)
		}
	}

	if err != nil {
		http.Error(w, "failed to update order", http.StatusInternalServerError)
		return
	}

	return
}
