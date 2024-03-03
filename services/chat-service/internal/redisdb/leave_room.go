package redisdb

import "context"

func (cache *redisCache) LeaveRoom(userId string, roomId string) {
	var ctx = context.Background()
	client := cache.client

	client.SRem(ctx, getKeyUserRooms(userId), roomId)
	client.SRem(ctx, getKeyRoomMembers(roomId), userId)

	memberCount, _ := client.SCard(ctx, getKeyRoomMembers(roomId)).Result()

	if memberCount == 1 {
		cache.DeleteRoom(roomId)
	}
}
