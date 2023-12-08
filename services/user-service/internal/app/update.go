package app

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
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

	user, err := s.userRepository.GetUserById(userID)
	if err != nil {
		return nil, status.Error(codes.Internal, "user does not exist")
	}

	user.Name = request.Name
	user.Picture = request.Picture
	user.Bio = request.Bio
	user.Website = request.Website
	user.PhoneNumber = request.PhoneNumber
	user.Banner = request.Banner

	_, err = s.userRepository.UpdateUser(&user)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, "failed to update user")
	}

	user2, err := s.userRepository.GetUserById(userID)

	err = pubsub.PublishProtoWithBinaryEncoding("user-topic", &pb.UserEvent{
		Event: &pb.UserEvent_Update{
			Update: &pb.UpdateUser{
				User: user2.ToUserPb(),
			},
		},
	})

	return user2.ToUserPb(), err
}
