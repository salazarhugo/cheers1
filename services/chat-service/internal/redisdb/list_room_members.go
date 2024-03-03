package redisdb

import "context"

func (cache *redisCache) GetRoomMembers(
	roomId string,
) ([]string, error) {
	members, err := cache.client.SMembers(
		context.Background(),
		getKeyRoomMembers(roomId),
	).Result()
	if err != nil {
		return nil, err
	}

	return members, nil
}
