package app

import (
	"context"
	"github.com/labstack/gommon/log"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetUsersIn(
	ctx context.Context,
	request *pb.GetUsersInRequest,
) (*pb.GetUsersInResponse, error) {
	users, err := s.userRepository.GetUsersIn(request.GetUserIds())
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, "failed to get users in")
	}

	return &pb.GetUsersInResponse{Users: users}, nil
}
