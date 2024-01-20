package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/friendship/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) ListFriend(
	ctx context.Context,
	request *friendship.ListFriendRequest,
) (*friendship.ListFriendResponse, error) {
	viewerID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve userID")
	}

	users, err := s.friendshipRepository.ListFriend(
		viewerID,
		1,
		100,
		request.UserId,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &friendship.ListFriendResponse{
		Items: users,
	}, nil
}
