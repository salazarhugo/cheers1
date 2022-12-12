package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListGoing(
	ctx context.Context,
	request *party.ListGoingRequest,
) (*party.ListGoingResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving userID")
	}

	response, err := s.partyRepository.ListGoing(userID, request.PartyId)
	if err != nil {
		return nil, err
	}

	return &party.ListGoingResponse{Users: response}, nil
}
