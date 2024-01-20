package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/friendship/v1"
)

func (s *Server) ListFriendIds(
	ctx context.Context,
	request *friendship.ListFriendIdsRequest,
) (*friendship.ListFriendIdsResponse, error) {
	friends, err := s.friendshipRepository.ListFriend("", 1, 100, request.UserId)
	if err != nil {
		return nil, err
	}

	friendIds := make([]string, 0)
	for _, u := range friends {
		friendIds = append(friendIds, u.Id)
	}

	return &friendship.ListFriendIdsResponse{
		Ids: friendIds,
	}, nil
}
