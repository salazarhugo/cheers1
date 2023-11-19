package repository

import (
	"github.com/google/uuid"
	"github.com/salazarhugo/cheers1/services/party-service/internal/model"
)

func (p *partyRepository) CreateParty(
	userID string,
	party *model.Party,
) (string, error) {
	db := p.spanner

	party.ID = uuid.NewString()
	party.UserID = userID

	result := db.Create(&party)
	if result.Error != nil {
		return "", result.Error
	}

	return party.ID, nil
}
