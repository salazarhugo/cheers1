package redisdb

import (
	"context"
	"strconv"
)

func (cache *redisCache) GetUnreadCount(
	roomID string,
	userID string,
) (int64, error) {
	ctx := context.Background()

	res, err := cache.client.Get(ctx, getKeyUnreadCount(roomID, userID)).Result()
	if err != nil {
		return 0, err
	}

	count, err := strconv.ParseInt(res, 10, 64)
	if err != nil {
		return 0, err
	}

	return count, nil
}
