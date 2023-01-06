package repository

import (
	"context"
	"fmt"
)

func getKeyFriendRequests(userId string) string {
	return fmt.Sprintf("%s:%s", keyFriendRequests, userId)
}

func getKeyFriends(userId string) string {
	return fmt.Sprintf("%s:%s", keyFriends, userId)
}

func (r repository) CreateFriendRequest(
	userId string,
	friendId string,
) error {
	err := r.redis.SAdd(
		context.Background(),
		getKeyFriendRequests(friendId),
		userId,
	).Err()

	return err
}
