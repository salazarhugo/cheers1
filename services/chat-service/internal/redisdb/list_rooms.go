package redisdb

import (
	"github.com/salazarhugo/cheers1/services/chat-service/internal/models"
)

func (cache *redisCache) ListRooms(userId string) ([]*models.Chat, error) {
	chatIDs, err := cache.ListChatIds(userId)
	if err != nil {
		return nil, err
	}

	var chats []*models.Chat
	for _, chatID := range chatIDs {
		room, err := cache.GetRoomWithId(userId, chatID)
		if err != nil {
			return nil, err
		}

		chats = append(chats, room)
	}

	return chats, nil
}
