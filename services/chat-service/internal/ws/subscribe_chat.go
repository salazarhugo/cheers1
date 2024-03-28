package ws

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v9"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/models"
	"log"
)

func SubscribeChat(
	conn *Connection,
	client *redis.Client,
	chatID string,
	viewerID string,
) {
	sub := client.Subscribe(context.Background(), chatID)
	ch := sub.Channel()
	for {
		select {
		case msg := <-ch:
			// Received msg from redis pub-sub
			var w models.WebSocketMessage
			err := json.Unmarshal([]byte(msg.Payload), &w)

			w.IsViewer = viewerID == w.UserId

			bytes, err := json.Marshal(&w)

			err = conn.Send(bytes)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
