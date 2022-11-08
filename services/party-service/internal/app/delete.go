package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
)

func (s *Server) DeleteParty(
	ctx context.Context,
	request *pb.DeletePartyRequest,
) (*pb.DeletePartyResponse, error) {
	return &pb.DeletePartyResponse{}, nil
}
