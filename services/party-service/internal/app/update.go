package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/party-service/internal/repository"
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

	err = AuthorizeUpdatePartyRequest(ctx, s.partyRepository, request)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = s.partyRepository.UpdateParty(request.Party)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to update party")
	}

	party, err := s.partyRepository.GetParty(request.Party.Id)
	if err != nil {
		return nil, err
	}

	return &pb.UpdatePartyResponse{
		Party: party,
	}, nil
}

func AuthorizeUpdatePartyRequest(
	ctx context.Context,
	repository repository.PartyRepository,
	request *pb.UpdatePartyRequest,
) error {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return status.Error(codes.Internal, "failed retrieving userID")
	}

	party, err := repository.GetParty(request.Party.Id)
	if err != nil {
		return err
	}

	app := utils.InitializeAppDefault()
	client, err := app.Auth(ctx)
	if err != nil {
		return err
	}

	log.Println(userID)
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
	if party.HostId != userID {
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

	if party.Id == "" {
		return status.Error(codes.InvalidArgument, "missing party id")
	}

	return nil
}
