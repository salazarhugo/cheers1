package app

import (
	"context"
	"github.com/labstack/gommon/log"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetUserItemsIn(
	ctx context.Context,
	request *pb.GetUserItemsInRequest,
) (*pb.GetUserItemsInResponse, error) {
	userID, err := GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed retrieving userID")
	}

	users, err := s.userRepository.GetUserItemsIn(userID, request.GetUserIds())
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, "failed to get users in")
	}

	return &pb.GetUserItemsInResponse{Users: users}, nil
}
