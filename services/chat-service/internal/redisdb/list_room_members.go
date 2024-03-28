package redisdb

import "context"

func (cache *redisCache) GetRoomMembers(
	chatID string,
) ([]string, error) {
	members, err := cache.client.SMembers(
		context.Background(),
		getKeyRoomMembers(chatID),
	).Result()
	if err != nil {
		return nil, err
	}

	return members, nil
}
