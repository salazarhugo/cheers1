package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/friendship/v1"
	"log"
)

func (s *Server) ListFriend(
	ctx context.Context,
	request *friendship.ListFriendRequest,
) (*friendship.ListFriendResponse, error) {
	items, err := s.friendshipRepository.ListFriend(request.UserId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &friendship.ListFriendResponse{
		Items: items,
	}, nil
}
