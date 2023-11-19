package repository

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	"github.com/salazarhugo/cheers1/services/party-service/internal/model"
)

func (p *partyRepository) GetPartyItem(
	userID string,
	partyID string,
) (*pb.PartyItem, error) {
	db := p.spanner

	var party model.Party

	result := db.Where("id = ?", partyID).First(&party)
	if result.Error != nil {
		return nil, result.Error
	}

	item := &pb.PartyItem{
		Party:             party.ToPartyPb(),
		User:              nil,
		GoingCount:        0,
		InterestedCount:   0,
		InvitedCount:      0,
		IsCreator:         false,
		ViewerWatchStatus: 0,
		MutualGoing:       nil,
		MutualInterested:  nil,
	}

	return item, nil
}
