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
	chatID := getDirectRoomId(userId, otherUserId)

	exists, err := client.Exists(context.Background(), getKeyRoom(chatID)).Result()
	if err != nil {
		return nil, err
	}

	if exists == 1 {
		room, err := cache.GetChat(userId, chatID)
		if err != nil {
			return nil, err
		}
		return room, nil
	}

	if exists == 0 {
		room := &models.Chat{
			Id:              chatID,
			LastMessageTime: 0,
			Status:          pb.RoomStatus_EMPTY.String(),
			Type:            models.ChatType_DIRECT,
			Admins:          []string{userId, otherUserId},
		}

		err := cache.CreateRoom(chatID, room)
		if err != nil {
			return nil, err
		}

		client.SAdd(ctx, getKeyUserRooms(userId), chatID)
		client.SAdd(ctx, getKeyUserRooms(otherUserId), chatID)
		client.SAdd(ctx, getKeyRoomMembers(room.Id), userId, otherUserId)

		err = cache.Publish(ctx, userId, room.Id).Err()
		err = cache.Publish(ctx, otherUserId, room.Id).Err()

		room_, err := cache.GetChat(userId, chatID)
		if err != nil {
			return nil, err
		}

		return room_, nil
	}

	return nil, nil
}
