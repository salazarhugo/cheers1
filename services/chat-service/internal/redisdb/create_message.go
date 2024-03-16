package redisdb

import (
	"context"
	json2 "encoding/json"
	"github.com/go-redis/redis/v9"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/models"
	"time"
)

func (cache *redisCache) CreateMessage(
	msg *models.ChatMessage,
) error {
	client := cache.client
	ctx := context.Background()

	bytes, err := json2.Marshal(msg)
	if err != nil {
		return err
	}

	err = client.Watch(ctx, func(tx *redis.Tx) error {
		cmd := tx.ZAdd(ctx, getKeyRoomMessages(msg.ChatId), redis.Z{
			Score:  float64(msg.CreatedAt),
			Member: string(bytes),
		})

		if cmd.Err() != nil {
			return cmd.Err()
		}

		cmd2 := tx.Expire(context.Background(), getKeyRoomMessages(msg.ChatId), 24*time.Hour)
		if cmd2.Err() != nil {
			return cmd.Err()
		}

		return nil
	})

	return err
}
