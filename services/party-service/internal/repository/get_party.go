package repository

import (
	"github.com/salazarhugo/cheers1/services/party-service/internal/model"
)

func (p *partyRepository) GetPartyById(
	id string,
) (*model.Party, error) {
	db := p.spanner
	var party model.Party

	result := db.Where("id = ?", id).First(&party)
	if result.Error != nil {
		return nil, result.Error
	}

	return &party, nil
}
