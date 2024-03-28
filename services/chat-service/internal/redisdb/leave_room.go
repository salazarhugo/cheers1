package redisdb

import "context"

func (cache *redisCache) LeaveChat(userId string, chatID string) {
	var ctx = context.Background()
	client := cache.client

	client.SRem(ctx, getKeyUserRooms(userId), chatID)
	client.SRem(ctx, getKeyRoomMembers(chatID), userId)

	memberCount, _ := client.SCard(ctx, getKeyRoomMembers(chatID)).Result()

	if memberCount == 1 {
		cache.DeleteRoom(chatID)
	}
}
