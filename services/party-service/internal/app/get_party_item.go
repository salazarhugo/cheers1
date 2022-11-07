package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetPartyItem(
	ctx context.Context,
	request *pb.GetPartyItemRequest,
) (*pb.GetPartyItemResponse, error) {
	err := ValidateGetPartyItemRequest(request)
	if err != nil {
		return nil, err
	}

	item, err := s.partyRepository.GetPartyItem(request.GetPartyId())
	if err != nil {
		return nil, err
	}

	return &pb.GetPartyItemResponse{
		Item: item,
	}, nil
}

func ValidateGetPartyItemRequest(request *pb.GetPartyItemRequest) error {
	partyId := request.GetPartyId()
	if partyId == "" {
		return status.Error(codes.InvalidArgument, "party_id parameter can't be nil")
	}

	return nil
}
