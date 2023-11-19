package repository

import (
	"github.com/salazarhugo/cheers1/services/party-service/internal/model"
)

func (p *partyRepository) UpdateParty(
	party *model.Party,
) (string, error) {
	db := p.spanner

	result := db.Updates(&party)
	if result.Error != nil {
		return "", result.Error
	}

	return party.ID, nil
}
