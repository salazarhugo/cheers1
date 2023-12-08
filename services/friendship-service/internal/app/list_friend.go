package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/friendship/v1"
	"github.com/salazarhugo/cheers1/services/friendship-service/internal/service"
	"log"
)

func (s *Server) ListFriend(
	ctx context.Context,
	request *friendship.ListFriendRequest,
) (*friendship.ListFriendResponse, error) {
	friendIds, err := s.friendshipRepository.ListFriend(request.UserId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	users, err := service.GetUsers(friendIds)
	if err != nil {
		return nil, err
	}

	return &friendship.ListFriendResponse{
		Items: users,
	}, nil
}
