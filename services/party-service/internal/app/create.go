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

func (s *Server) CreateParty(
	ctx context.Context,
	request *pb.CreatePartyRequest,
) (*pb.CreatePartyResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed retrieving userID")
	}

	partyReq := request.GetParty()
	log.Println(partyReq)
	if partyReq == nil {
		return nil, status.Error(codes.InvalidArgument, "party parameter can't be nil")
	}

	partyReq.HostId = userID

	party := model.ToPartyDomain(partyReq)

	partyID, err := s.partyRepository.CreateParty(userID, party)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create party")
	}

	response, err := s.partyRepository.GetPartyById(partyID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get new party")
	}

	return &pb.CreatePartyResponse{
		Party: response.ToPartyPb(),
	}, nil
}
