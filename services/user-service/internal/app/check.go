package app

import (
	"context"
	"github.com/labstack/gommon/log"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CheckUsername(
	ctx context.Context,
	request *pb.CheckUsernameRequest,
) (*pb.CheckUsernameResponse, error) {
	_, err := GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "not signed in")
	}

	exists, err := s.userRepository.CheckUsername(request.Username)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, "failed to check username")
	}

	return &pb.CheckUsernameResponse{
		Valid: !exists,
	}, nil
}
