package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/friendship/v1"
)

func (s *Server) DeleteFriendship(
	ctx context.Context,
	request *friendship.DeletedFriendship,
) (*friendship.DeletedFriendship, error) {
	return &friendship.DeletedFriendship{}, nil
}
