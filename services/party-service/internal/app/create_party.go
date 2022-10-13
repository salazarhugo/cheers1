package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/party/v1"
	party "github.com/salazarhugo/cheers1/genproto/cheers/type/party"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateParty(
	ctx context.Context,
	request *pb.CreatePartyRequest,
) (*party.Party, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving userID")
	}

	partyReq := request.GetParty()
	if partyReq == nil {
		return nil, status.Error(codes.InvalidArgument, "party parameter can't be nil")
	}

	partyReq.HostId = userID

	response, err := s.partyRepository.CreateParty(userID, partyReq)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create party")
	}

	return response, nil
}
