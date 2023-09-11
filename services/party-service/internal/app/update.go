package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	partypb "github.com/salazarhugo/cheers1/gen/go/cheers/type/party"
	"github.com/salazarhugo/cheers1/libs/utils"
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

	var party *partypb.Party

	if request.Party.Name != "" {
		party, err = s.partyRepository.GetPartyWithName(request.Party.GetName())
	} else {
		party, err = s.partyRepository.GetParty(request.Party.Id)
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = AuthorizeUpdatePartyRequest(ctx, party)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	request.Party.Id = party.Id
	_, err = s.partyRepository.UpdateParty(request.Party)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to update party")
	}

	updatedParty, err := s.partyRepository.GetParty(request.Party.Id)
	if err != nil {
		return nil, err
	}

	return &pb.UpdatePartyResponse{
		Party: updatedParty,
	}, nil
}

func AuthorizeUpdatePartyRequest(
	ctx context.Context,
	party *partypb.Party,
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

	if party.Id == "" && party.Name == "" {
		return status.Error(codes.InvalidArgument, "party id or name required")
	}

	return nil
}
