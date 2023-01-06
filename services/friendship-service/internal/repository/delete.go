package repository

import (
	"context"
)

func (r repository) DeleteFriendship(
	postId string,
	friendshipId string,
) error {
	return r.redis.ZRem(context.Background(), getKeyFriendRequests(postId), friendshipId).Err()
}
