package redisdb

import (
	"context"
	json2 "encoding/json"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (cache *redisCache) GetRoomWithId(
	userId string,
	roomId string,
) (*models.Chat, error) {
	client := cache.client

	val, err := client.Get(
		context.Background(),
		getKeyRoom(roomId),
	).Result()
	if err == redis.Nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Room %s not found", roomId))
	}

	if err != nil {
		return nil, err
	}

	room := models.Chat{}
	err = json2.Unmarshal([]byte(val), &room)
	if err != nil {
		return nil, err
	}

	otherUserId, err := cache.GetOtherUserId(roomId, userId)
	lastMessage, err := getLatestMessage(client, roomId)
	unreadCount, err := cache.GetUnreadCount(roomId, userId)

	if lastMessage != nil {
		room.LastMessageText = lastMessage.Text
		room.LastMessageTime = lastMessage.CreateTime
		room.LastMessageType = lastMessage.Type.String()
	}

	room.UnreadCount = unreadCount

	status, err := cache.GetRoomStatus(roomId, userId, otherUserId)
	room.Status = status.String()

	return &room, nil
}
