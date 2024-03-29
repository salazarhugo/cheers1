package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/friendship/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) DeleteFriend(
	ctx context.Context,
	request *friendship.DeleteFriendRequest2,
) (*friendship.DeleteFriendResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve userID")
	}

	err = s.friendshipRepository.DeleteFriend(
		userID,
		request.UserId,
	)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &friendship.DeleteFriendResponse{}, nil
}
