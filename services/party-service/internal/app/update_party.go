package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
)

func (s *Server) UpdateParty(
	ctx context.Context,
	request *pb.UpdatePartyRequest,
) (*pb.UpdatePartyResponse, error) {
	return &pb.UpdatePartyResponse{}, nil
}
