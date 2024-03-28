package redisdb

import "context"

func (cache *redisCache) IsMember(
	userId string,
	chatID string,
) (bool, error) {
	isMember, err := cache.client.SIsMember(
		context.Background(),
		getKeyRoomMembers(chatID),
		userId,
	).Result()

	if err != nil {
		return false, err
	}

	return isMember, nil
}
