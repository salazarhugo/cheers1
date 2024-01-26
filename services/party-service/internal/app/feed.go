package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
)

func (s *Server) FeedParty(
	ctx context.Context,
	request *pb.FeedPartyRequest,
) (*pb.FeedPartyResponse, error) {
	userID, _ := utils.GetUserId(ctx)

	response, err := s.partyRepository.FeedParty(userID, request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
