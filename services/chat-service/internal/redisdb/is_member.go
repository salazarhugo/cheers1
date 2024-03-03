package redisdb

import "context"

func (cache *redisCache) IsMember(
	userId string,
	roomId string,
) (bool, error) {
	isMember, err := cache.client.SIsMember(
		context.Background(),
		getKeyRoomMembers(roomId),
		userId,
	).Result()

	if err != nil {
		return false, err
	}

	return isMember, nil
}
