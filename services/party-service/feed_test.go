package user_service

import (
	party3 "github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	"github.com/salazarhugo/cheers1/services/party-service/internal/repository"
	"log"
	"testing"
)

func TestFeedParty(t *testing.T) {
	// Create a mock repository
	repo := repository.NewRepository()

	// This user must exist in the database
	userID := "user1"

	request := &party3.FeedPartyRequest{
		Parent:   "",
		PageSize: 10,
		Page:     1,
	}

	parties, err := repo.FeedParty(userID, request)
	if err != nil {
		t.Error(err)
		return
	}

	log.Println(parties)
}
