package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateParty(
	ctx context.Context,
	request *pb.UpdatePartyRequest,
) (*pb.UpdatePartyResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed retrieving userID")
	}

	err = ValidateUpdatePartyRequest(request)
	if err != nil {
		return nil, err
	}

	partyReq := request.GetParty()
	partyReq.HostId = userID

	partyID, err := s.partyRepository.UpdateParty(partyReq)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create party")
	}

	party, err := s.partyRepository.GetParty(partyID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get party")
	}

	return &pb.UpdatePartyResponse{
		Party: party,
	}, nil
}

func ValidateUpdatePartyRequest(request *pb.UpdatePartyRequest) error {
	if request == nil {
		return status.Error(codes.InvalidArgument, "party parameter can't be nil")
	}

	party := request.GetParty()

	if party.Id == "" {
		return status.Error(codes.InvalidArgument, "missing party id")
	}

	return nil
}
