package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/friendship/v1"
	"log"
)

func (s *Server) ListFriendIds(
	ctx context.Context,
	request *friendship.ListFriendIdsRequest,
) (*friendship.ListFriendIdsResponse, error) {
	friendIds, err := s.friendshipRepository.ListFriend(request.UserId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &friendship.ListFriendIdsResponse{
		Ids: friendIds,
	}, nil
}
