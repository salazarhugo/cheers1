package app

import (
	"context"
	"github.com/labstack/gommon/log"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListFollowers(
	ctx context.Context,
	request *pb.ListFollowersRequest,
) (*pb.ListFollowersResponse, error) {
	userID, err := GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving userID")
	}

	followers, err := s.userRepository.ListFollowers(userID, request)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &pb.ListFollowersResponse{
		Users: followers,
	}, nil
}
