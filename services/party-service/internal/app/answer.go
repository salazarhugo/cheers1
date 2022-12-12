package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) AnswerParty(
	ctx context.Context,
	request *pb.AnswerPartyRequest,
) (*pb.AnswerPartyResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed retrieving userID")
	}

	switch request.Answer {
	case pb.PartyAnswer_GOING:
		err = s.partyRepository.GoingParty(userID, request.PartyId)
	case pb.PartyAnswer_NOT_INTERESTED:
		err = s.partyRepository.UnGoingParty(userID, request.PartyId)
	}

	if err != nil {
		return nil, status.Error(codes.Internal, "failed to answer party")
	}

	return &pb.AnswerPartyResponse{}, nil
}
