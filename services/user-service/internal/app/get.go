package app

import (
	"context"
	"github.com/labstack/gommon/log"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetUser(
	ctx context.Context,
	request *pb.GetUserRequest,
) (*pb.GetUserResponse, error) {
	userID, err := GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed retrieving userID")
	}

	otherUserID := request.GetUserId()

	response, err := s.userRepository.GetUserWithViewer(userID, otherUserID)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return response, nil
}
