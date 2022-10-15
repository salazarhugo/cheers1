package app

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/salazarhugo/cheers1/genproto/cheers/type/user"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUser(
	ctx context.Context,
	request *pb.CreateUserRequest,
) (*user.User, error) {
	userID, err := GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving userID")
	}

	partyReq := request.GetUser()
	if partyReq == nil {
		return nil, status.Error(codes.InvalidArgument, "post parameter can't be nil")
	}

	postID, err := s.userRepository.CreateUser(userID, partyReq)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, "failed to create user")
	}

	user, err := s.userRepository.GetUser(userID, postID)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, "failed to get user")
	}

	return user.User, nil
}
