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

func (cache *redisCache) GetChat(
	viewerID string,
	chatID string,
) (*models.Chat, error) {
	client := cache.client

	val, err := client.Get(
		context.Background(),
		getKeyRoom(chatID),
	).Result()
	if err == redis.Nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Room %s not found", chatID))
	}

	if err != nil {
		return nil, err
	}

	chat := models.Chat{}
	err = json2.Unmarshal([]byte(val), &chat)
	if err != nil {
		return nil, err
	}

	otherUserId, err := cache.GetOtherUserId(chatID, viewerID)
	lastMessage, err := getLatestMessage(client, chatID)
	unreadCount, err := cache.GetUnreadCount(chatID, viewerID)
	lastReadOfViewer, err := cache.GetLastRead(chatID, viewerID)
	lastReadOfOtherUser, err := cache.GetLastRead(chatID, otherUserId)

	if lastMessage != nil {
		chat.LastMessageText = lastMessage.Text
		chat.LastMessageTime = lastMessage.CreatedAt
		chat.LastMessageType = "TEXT"
	}

	chat.Status = GetRoomStatus(
		viewerID,
		lastMessage,
		lastReadOfViewer,
		lastReadOfOtherUser,
	)

	chat.UnreadCount = unreadCount

	return &chat, nil
}

func GetRoomStatus(
	viewerID string,
	lastMessage *models.ChatMessage,
	lastReadOfViewer int64,
	lastReadOfOtherUser int64,
) string {
	if lastMessage == nil {
		return "EMPTY"
	}

	isLastMessageFromViewer := lastMessage.UserId == viewerID

	if isLastMessageFromViewer {
		if lastReadOfOtherUser > lastMessage.CreatedAt {
			return "OPENED"
		}
		return "SENT"
	}

	if lastReadOfViewer > lastMessage.CreatedAt {
		return "RECEIVED"
	}

	return "NEW"
}
