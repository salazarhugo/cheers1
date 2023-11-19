package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetParty(
	ctx context.Context,
	request *pb.GetPartyRequest,
) (*pb.GetPartyResponse, error) {
	err := ValidateGetPartyRequest(request)
	if err != nil {
		return nil, err
	}

	party, err := s.partyRepository.GetPartyById(request.GetPartyId())
	if err != nil {
		return nil, err
	}

	return &pb.GetPartyResponse{
		Party: party.ToPartyPb(),
	}, nil
}

func ValidateGetPartyRequest(request *pb.GetPartyRequest) error {
	partyId := request.GetPartyId()
	if partyId == "" {
		return status.Error(codes.InvalidArgument, "party_id parameter can't be nil")
	}

	return nil
}
