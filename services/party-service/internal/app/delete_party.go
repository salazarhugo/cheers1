package app

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/party/v1"
)

func (s *Server) DeleteParty(
	ctx context.Context,
	request *pb.DeletePartyRequest,
) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
