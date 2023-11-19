package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/party-service/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) UpdateParty(
	ctx context.Context,
	request *pb.UpdatePartyRequest,
) (*pb.UpdatePartyResponse, error) {
	err := ValidateUpdatePartyRequest(request)
	if err != nil {
		return nil, err
	}

	party, err := s.partyRepository.GetPartyById(request.Party.Id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = AuthorizeUpdatePartyRequest(ctx, party)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	request.Party.Id = party.ID
	updatedPartyRequest := model.ToPartyDomain(request.Party)
	partyID, err := s.partyRepository.UpdateParty(updatedPartyRequest)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to update party")
	}

	updatedParty, err := s.partyRepository.GetPartyById(partyID)
	if err != nil {
		return nil, err
	}

	return &pb.UpdatePartyResponse{
		Party: updatedParty.ToPartyPb(),
	}, nil
}

func AuthorizeUpdatePartyRequest(
	ctx context.Context,
	party *model.Party,
) error {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return status.Error(codes.Internal, "failed retrieving userID")
	}

	if err != nil {
		return err
	}

	app := utils.InitializeAppDefault()
	client, err := app.Auth(ctx)
	if err != nil {
		return err
	}

	user, err := client.GetUser(ctx, userID)
	if err != nil {
		log.Println(err)
		return err
	}

	isAdmin := user.CustomClaims["admin"] != nil

	if isAdmin == true {
		return nil
	}

	// Check if user is the party host
	if party.UserID != userID {
		return status.Error(codes.PermissionDenied, "insufficient permissions")
	}

	return nil
}

func ValidateUpdatePartyRequest(request *pb.UpdatePartyRequest) error {
	if request == nil {
		return status.Error(codes.InvalidArgument, "request can't be nil")
	}

	party := request.GetParty()

	if request.Party == nil {
		return status.Error(codes.InvalidArgument, "party parameter can't be nil")
	}

	if party.Id == "" && party.Name == "" {
		return status.Error(codes.InvalidArgument, "party id or name required")
	}

	return nil
}
