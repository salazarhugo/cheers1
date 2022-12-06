package app

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUser(
	ctx context.Context,
	request *pb.CreateUserRequest,
) (*pb.CreateUserResponse, error) {
	userID, err := GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving userID")
	}

	userID, err = s.userRepository.CreateUser(userID, &user.User{
		Username: request.Username,
		Name:     request.Name,
		Picture:  request.Picture,
		Email:    request.Email,
	})
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, "failed to create user")
	}

	response, err := s.userRepository.GetUser(userID, userID)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, "failed to get user")
	}

	return &pb.CreateUserResponse{
		User: response.User,
	}, nil
}
