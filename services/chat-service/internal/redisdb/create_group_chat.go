package redisdb

import (
	"context"
	"github.com/google/uuid"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/models"
)

func (cache *redisCache) CreateGroup(
	name string,
	UUIDs []string,
) (*models.Chat, error) {
	client := cache.client

	ctx := context.Background()

	room := &models.Chat{
		Id:              uuid.New().String(),
		LastMessageTime: 0,
		Name:            name,
		Status:          "",
		Type:            models.ChatType_GROUP,
		Admins:          []string{UUIDs[0]},
	}

	err := cache.CreateRoom(room.Id, room)
	if err != nil {
		return nil, err
	}

	for _, uid := range UUIDs {
		client.SAdd(ctx, getKeyUserRooms(uid), room.Id)
		client.SAdd(ctx, getKeyRoomMembers(room.Id), uid)
		err = cache.Publish(ctx, uid, room.Id).Err()
	}

	return room, nil
}
