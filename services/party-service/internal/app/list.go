package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListParty(
	ctx context.Context,
	request *pb.ListPartyRequest,
) (*pb.ListPartyResponse, error) {
	viewerID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving userID")
	}

	response, err := s.partyRepository.ListParty(viewerID, request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
