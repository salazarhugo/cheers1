package app

import (
	"context"
	"github.com/labstack/gommon/log"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListFollowing(
	ctx context.Context,
	request *pb.ListFollowingRequest,
) (*pb.ListFollowingResponse, error) {
	userID, err := GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed retrieving userID")
	}

	following, err := s.userRepository.ListFollowing(userID, request)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &pb.ListFollowingResponse{
		Users: following,
	}, nil
}
