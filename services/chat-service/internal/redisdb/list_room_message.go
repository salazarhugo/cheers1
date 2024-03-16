package redisdb

import (
	"context"
	json2 "encoding/json"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/models"
)

func (cache *redisCache) ListChatMessage(
	roomId string,
	offset int,
	limit int,
) ([]*models.ChatMessage, error) {
	client := cache.client

	values, err := client.ZRevRange(
		context.Background(),
		getKeyRoomMessages(roomId),
		int64(offset),
		int64(offset+limit),
	).Result()
	if err != nil {
		panic(err)
	}
	var messages []*models.ChatMessage

	for _, msg := range values {
		chatMsg := &models.ChatMessage{}
		err := json2.Unmarshal([]byte(msg), chatMsg)
		if err != nil {
			return nil, err
		}
		//chatMsg.IsSender =
		messages = append(messages, chatMsg)
	}

	return messages, nil
}
