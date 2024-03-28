package redisdb

import "context"

func (cache *redisCache) DeleteRoom(chatID string) error {
	var ctx = context.Background()
	client := cache.client

	members, err := client.SMembers(ctx, getKeyRoomMembers(chatID)).Result()
	if err != nil {
		return err
	}

	for _, uid := range members {
		// Delete user rooms
		client.SRem(ctx, getKeyUserRooms(uid), chatID)
		// Delete user last read timestamps
		client.Del(ctx, getKeyLastRead(chatID, uid))
	}

	// Delete room
	client.Del(ctx, getKeyRoom(chatID))
	// Delete room seen
	client.Del(ctx, getKeyRoomSeen(chatID))
	// Delete room members
	client.Del(ctx, getKeyRoomMembers(chatID))
	// Delete room messages
	client.Del(ctx, getKeyRoomMessages(chatID))

	return nil
}
