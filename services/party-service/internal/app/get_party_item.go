package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetPartyItem(
	ctx context.Context,
	request *pb.GetPartyItemRequest,
) (*pb.GetPartyItemResponse, error) {
	userID, _ := utils.GetUserId(ctx)

	err := ValidateGetPartyItemRequest(request)
	if err != nil {
		return nil, err
	}

	var item *pb.PartyItem
	if userID == "" {
		item, err = s.partyRepository.GetPartyItemPublic(request.PartyId)
	} else {
		item, err = s.partyRepository.GetPartyItem(userID, request.PartyId)
	}

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
