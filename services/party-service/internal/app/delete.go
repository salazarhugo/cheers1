package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteParty(
	ctx context.Context,
	request *pb.DeletePartyRequest,
) (*pb.DeletePartyResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed retrieving userID")
	}

	if request.PartyId == "" {
		return nil, status.Error(codes.InvalidArgument, "parameter can't be blank")
	}

	app := utils.InitializeAppDefault()
	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}
	user, err := client.GetUser(ctx, userID)
	isAdmin := user.CustomClaims["admin"] != nil

	party, err := s.partyRepository.GetPartyById(request.PartyId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "party not found")
	}

	if party.UserID != userID && !isAdmin {
		return nil, status.Error(codes.PermissionDenied, "insufficient permissions: must be creator of the party or admin")
	}

	err = s.partyRepository.DeletePartyById(request.PartyId)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to delete party")
	}

	return &pb.DeletePartyResponse{}, nil
}
