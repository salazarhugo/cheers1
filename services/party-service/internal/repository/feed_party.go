package repository

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	"github.com/salazarhugo/cheers1/services/party-service/internal/model"
	"time"
)

func (p *partyRepository) FeedParty(
	userID string,
	request *pb.FeedPartyRequest,
) (*pb.FeedPartyResponse, error) {
	db := p.spanner

	page := request.GetPage()
	if page == 0 {
		page = 1
	}
	pageSize := request.GetPageSize()
	if pageSize == 0 {
		pageSize = 18
	}
	skip := pageSize * (page - 1)

	var parties []model.Party

	result := db.Limit(int(pageSize)).Offset(int(skip)).Where("end_date > ?", time.Now()).Order("end_date asc").Find(&parties)
	if result.Error != nil {
		return nil, result.Error
	}

	items := make([]*pb.PartyItem, 0)
	for _, party := range parties {
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
		items = append(items, item)
	}

	return &pb.FeedPartyResponse{
		Items:         items,
		NextPageToken: "token123",
	}, nil
}
