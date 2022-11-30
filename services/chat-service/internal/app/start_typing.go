package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
)

func (s *Server) TypingStart(ctx context.Context, req *pb.TypingReq) (*pb.Empty, error) {
	//TODO implement me
	return &pb.Empty{}, nil
}
