package redisdb

import "context"

func (cache *redisCache) SetUnreadCount(
	roomID string,
	userID string,
	count int64,
) error {
	ctx := context.Background()

	err := cache.client.Set(ctx, getKeyLastRead(roomID, userID), count, 0).Err()

	return err
}
