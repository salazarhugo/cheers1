package redisdb

import "context"

func (cache *redisCache) DeleteRoom(roomId string) error {
	var ctx = context.Background()
	client := cache.client

	members, err := client.SMembers(ctx, getKeyRoomMembers(roomId)).Result()
	if err != nil {
		return err
	}

	for _, uid := range members {
		// Delete user rooms
		client.SRem(ctx, getKeyUserRooms(uid), roomId)
		// Delete user last read timestamps
		client.Del(ctx, getKeyLastRead(roomId, uid))
	}

	// Delete room
	client.Del(ctx, getKeyRoom(roomId))
	// Delete room seen
	client.Del(ctx, getKeyRoomSeen(roomId))
	// Delete room members
	client.Del(ctx, getKeyRoomMembers(roomId))
	// Delete room messages
	client.Del(ctx, getKeyRoomMessages(roomId))

	return nil
}
