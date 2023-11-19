package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) TransferParty(
	ctx context.Context,
	request *pb.TransferPartyRequest,
) (*pb.TransferPartyResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed retrieving userID")
	}

	if request.PartyId == "" || request.UserId == "" {
		return nil, status.Error(codes.InvalidArgument, "parameter can't be blank")
	}

	app := utils.InitializeAppDefault()
	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	// Check if user is admin
	user, err := client.GetUser(ctx, userID)
	if user.CustomClaims["admin"] == nil {
		return nil, status.Error(codes.PermissionDenied, "insufficient permissions")
	}

	// Check if party exists
	party, err := s.partyRepository.GetPartyById(request.PartyId)
	if err != nil {
		return nil, err
	}

	err = s.partyRepository.TransferParty(
		request.UserId,
		party.ID,
	)

	if err != nil {
		return nil, status.Error(codes.Internal, "failed to transfer party")
	}

	return &pb.TransferPartyResponse{}, nil
}
