package redisdb

import "context"

func (cache *redisCache) ListChatIds(
	userID string,
) ([]string, error) {
	client := cache.client

	values, err := client.SMembers(
		context.Background(),
		getKeyUserRooms(userID),
	).Result()

	if err != nil {
		return nil, err
	}

	return values, nil
}
