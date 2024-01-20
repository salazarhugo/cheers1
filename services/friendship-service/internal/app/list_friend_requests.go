package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/friendship/v1"
	"log"
)

func (s *Server) ListFriendRequests(
	ctx context.Context,
	request *friendship.ListFriendRequestsRequest,
) (*friendship.ListFriendRequestsResponse, error) {

	items, err := s.friendshipRepository.ListFriendRequests(
		1,
		100,
		request.UserId,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &friendship.ListFriendRequestsResponse{
		Items: items,
	}, nil
}
