package app

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateUser(
	ctx context.Context,
	request *pb.UpdateUserRequest,
) (*user.User, error) {
	userID, err := GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed retrieving userID")
	}

	err = s.userRepository.UpdateUser(userID, request)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, "failed to update user")
	}

	result, err := s.userRepository.GetUser(userID, userID)

	return result.GetUser(), nil
}
