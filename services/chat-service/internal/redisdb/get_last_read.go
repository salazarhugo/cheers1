package redisdb

import (
	"context"
	"strconv"
)

func (cache *redisCache) GetLastRead(
	roomID string,
	userID string,
) (int64, error) {
	ctx := context.Background()

	res, err := cache.client.Get(ctx, getKeyLastRead(roomID, userID)).Result()
	if err != nil {
		return 0, err
	}

	lastReadAt, err := strconv.ParseInt(res, 10, 64)
	if err != nil {
		return 0, err
	}

	return lastReadAt, nil
}
