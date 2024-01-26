package user_service

import (
	party3 "github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	"github.com/salazarhugo/cheers1/services/party-service/internal/repository"
	"log"
	"testing"
)

func TestListParty(t *testing.T) {
	// Create a mock repository
	repo := repository.NewRepository()

	// This user must exist in the database
	userID := "Ii7mFI34JaYKYsOxYg90fYToa182"

	request := &party3.ListPartyRequest{
		UserId:   userID,
		PageSize: 10,
		Page:     1,
	}

	parties, err := repo.ListParty(userID, request)
	if err != nil {
		t.Error(err)
		return
	}

	log.Println(parties)
}
