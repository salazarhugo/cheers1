package user_service

import (
	model2 "github.com/salazarhugo/cheers1/services/party-service/internal/model"
	"github.com/salazarhugo/cheers1/services/party-service/internal/repository"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUpsertParty(t *testing.T) {
	// Create a mock repository
	repo := repository.NewRepository()

	// This user must exist in the database
	userID := "user1"

	model := &model2.Party{
		Slug:         "opium-club-1",
		UserID:       userID,
		Name:         "Opium Club 1",
		LocationName: "32 Avenue George V",
		Description:  "",
		Latitude:     0.45325,
		Longitude:    1.32532,
		StartDate:    time.Now(),
		EndDate:      time.Now().AddDate(0, 1, 5),
		City:         "Limassol",
	}

	partyID, err := repo.UpsertPartyBySlug(model)
	if err != nil {
		t.Error(err)
		return
	}

	party, err := repo.GetPartyById(partyID)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, model.Name, party.Name, "The two names should be the same.")
	assert.Equal(t, model.City, party.City, "The two cities should be the same.")

	model.Name = "Opium Club Updated"
	partyID2, err := repo.UpsertPartyBySlug(model)
	if err != nil {
		t.Error(err)
		return
	}

	party2, err := repo.GetPartyById(partyID2)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, model.Name, party2.Name, "The two names should be the same.")
	assert.Equal(t, party.ID, party2.ID, "Should be same id")

	err = repo.DeletePartyById(party.ID)
	if err != nil {
		t.Error(err)
		return
	}
}
