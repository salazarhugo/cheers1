package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DuplicateParty(
	ctx context.Context,
	request *pb.DuplicatePartyRequest,
) (*pb.DuplicatePartyResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed retrieving userID")
	}

	party, err := s.partyRepository.GetParty(request.PartyId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "can't get party")
	}

	party.HostId = userID
	party.Name = party.Name + " - Copy"

	_, err = s.partyRepository.CreateParty(userID, party)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create party")
	}

	return &pb.DuplicatePartyResponse{}, nil
}
