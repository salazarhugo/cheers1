package redisdb

import "context"

func (cache *redisCache) SetLastRead(
	roomID string,
	userID string,
	timestamp int64,
) error {
	ctx := context.Background()

	err := cache.client.Set(ctx, getKeyLastRead(roomID, userID), timestamp, 0).Err()

	return err
}
