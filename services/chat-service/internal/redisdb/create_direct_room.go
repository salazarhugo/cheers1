package redisdb

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/models"
)

func (cache *redisCache) GetOrCreateDirectRoom(
	userId string,
	otherUserId string,
) (*models.Chat, error) {
	client := cache.client

	ctx := context.Background()
	roomId := getDirectRoomId(userId, otherUserId)

	exists, err := client.Exists(context.Background(), getKeyRoom(roomId)).Result()
	if err != nil {
		return nil, err
	}

	if exists == 1 {
		room, err := cache.GetRoomWithId(userId, roomId)
		if err != nil {
			return nil, err
		}
		return room, nil
	}

	if exists == 0 {
		room := &models.Chat{
			Id:              roomId,
			LastMessageTime: 0,
			Status:          pb.RoomStatus_EMPTY.String(),
			Type:            models.ChatType_DIRECT,
			Admins:          []string{userId, otherUserId},
		}

		err := cache.CreateRoom(roomId, room)
		if err != nil {
			return nil, err
		}

		client.SAdd(ctx, getKeyUserRooms(userId), roomId)
		client.SAdd(ctx, getKeyUserRooms(otherUserId), roomId)
		client.SAdd(ctx, getKeyRoomMembers(room.Id), userId, otherUserId)

		room_, err := cache.GetRoomWithId(userId, roomId)
		if err != nil {
			return nil, err
		}

		return room_, nil
	}

	return nil, nil
}
