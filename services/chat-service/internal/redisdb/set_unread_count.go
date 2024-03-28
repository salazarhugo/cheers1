package redisdb

import "context"

func (cache *redisCache) ResetUnreadCount(
	roomID string,
	userID string,
) error {
	ctx := context.Background()

	err := cache.client.Set(ctx, getKeyUnreadCount(roomID, userID), 0, 0).Err()

	return err
}

func (cache *redisCache) IncrementUnreadCount(
	roomID string,
	userID string,
) error {
	ctx := context.Background()

	err := cache.client.Incr(ctx, getKeyUnreadCount(roomID, userID)).Err()

	return err
}
