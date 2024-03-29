package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/friendship/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) CreateFriendRequest(
	ctx context.Context,
	request *friendship.CreateFriendRequestRequest,
) (*friendship.CreateFriendRequestResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve userID")
	}

	isFriend, err := s.friendshipRepository.CheckFriend(
		userID,
		request.UserId,
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if isFriend {
		return nil, status.Error(codes.FailedPrecondition, "already friends")
	}

	err = s.friendshipRepository.CreateFriendRequest(
		userID,
		request.UserId,
	)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &friendship.CreateFriendRequestResponse{}, nil
}
