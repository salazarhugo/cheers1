package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	"github.com/salazarhugo/cheers1/services/party-service/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UpsertParty This method should only be called by scrape-parties cloud function
func (s *Server) UpsertParty(
	ctx context.Context,
	request *pb.UpsertPartyRequest,
) (*pb.UpsertPartyResponse, error) {
	partyReq := request.GetParty()
	if partyReq == nil {
		return nil, status.Error(codes.InvalidArgument, "party parameter can't be nil")
	}

	newParty := model.ToPartyDomain(partyReq)
	// official cheers userID
	newParty.UserID = "818dd7d8-55ae-40e1-b9c6-956691292e82"

	newPartyId, err := s.partyRepository.UpsertPartyBySlug(newParty)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to insert party")
	}

	response, err := s.partyRepository.GetPartyById(newPartyId)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get new party")
	}

	return &pb.UpsertPartyResponse{
		Party: response.ToPartyPb(),
	}, nil
}
