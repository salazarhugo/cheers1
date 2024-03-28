package redisdb

import (
	"context"
	json2 "encoding/json"
	"github.com/go-redis/redis/v9"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/models"
)

func getLatestMessage(
	client *redis.Client,
	chatID string,
) (*models.ChatMessage, error) {
	lastMessage, err := client.ZRevRangeByScore(
		context.Background(),
		getKeyRoomMessages(chatID),
		&redis.ZRangeBy{
			Min:    "-inf",
			Max:    "+inf",
			Offset: 0,
			Count:  1,
		}).Result()

	if len(lastMessage) == 0 {
		return nil, nil
	}

	message := &models.ChatMessage{}

	err = json2.Unmarshal([]byte(lastMessage[0]), message)
	if err != nil {
		return nil, err
	}

	return message, nil
}
