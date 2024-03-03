package redisdb

import (
	"context"
	json2 "encoding/json"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/models"
	"time"
)

func (cache *redisCache) CreateRoom(
	roomId string,
	room *models.Chat,
) error {
	client := cache.client

	json, err := json2.Marshal(room)
	if err != nil {
		return err
	}

	client.Set(
		context.Background(),
		getKeyRoom(roomId),
		json,
		cache.expires*time.Second,
	)

	return nil
}
