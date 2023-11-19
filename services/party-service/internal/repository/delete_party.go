package repository

import "github.com/salazarhugo/cheers1/services/party-service/internal/model"

func (p *partyRepository) DeletePartyById(
	id string,
) error {
	db := p.spanner
	result := db.Delete(&model.Party{ID: id})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
