package repository

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/party-service/internal/model"
)

func (p *partyRepository) ListParty(
	viewerID string,
	request *pb.ListPartyRequest,
) (*pb.ListPartyResponse, error) {
	spanner := p.spanner

	limit, offset := utils.GetLimitAndOffsetPagination(int(request.Page), int(request.PageSize))

	var parties []model.Party

	result := spanner.
		Limit(limit).
		Offset(offset).
		Where("user_id = ?", request.UserId).
		Order("end_date asc").
		Find(&parties)
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
			IsCreator:         viewerID == party.UserID,
			ViewerWatchStatus: 0,
			MutualGoing:       nil,
			MutualInterested:  nil,
		}
		items = append(items, item)
	}

	return &pb.ListPartyResponse{
		Items:         items,
		NextPageToken: "token123",
	}, nil
}
