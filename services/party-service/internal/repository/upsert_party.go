package repository

import "github.com/salazarhugo/cheers1/services/party-service/internal/model"

func (p *partyRepository) UpsertPartyBySlug(
	party *model.Party,
) (string, error) {
	oldParty, _ := p.GetPartyBySlug(party.Slug)
	exists := oldParty != nil
	if exists {
		party.ID = oldParty.ID
		return p.UpdateParty(party)
	} else {
		return p.CreateParty(party.UserID, party)
	}
}
