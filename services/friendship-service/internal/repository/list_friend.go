package repository

import (
	"context"
	"log"
)

func (r repository) ListFriend(userId string) ([]string, error) {
	ctx := context.Background()
	userIds, err := r.redis.SMembers(ctx, getKeyFriends(userId)).Result()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return userIds, err
}
