package app

import (
	"context"
	"github.com/labstack/gommon/log"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/utils/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) SearchUser(
	ctx context.Context,
	request *pb.SearchUserRequest,
) (*pb.SearchUserResponse, error) {
	_, err := GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving userID")
	}

	users, err := s.userRepository.SearchUser(request.GetQuery())

	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, "failed to search user")
	}

	return &pb.SearchUserResponse{
		Users: models.ToUserItemsPb(users),
	}, nil
}
