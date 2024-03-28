package ws

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v9"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/models"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/repository"
)

func SubscribeUser(
	conn *Connection,
	client *redis.Client,
	viewerID string,
) {
	sub := client.Subscribe(context.Background(), viewerID)
	ch := sub.Channel()

	for {
		select {
		case msg := <-ch:
			chatID := msg.Payload

			err := SendChatToWS(conn, client, viewerID, chatID)
			if err != nil {
				continue
			}

			go SubscribeChat(conn, client, chatID, viewerID)
		}
	}
}

func SendChatToWS(
	conn *Connection,
	client *redis.Client,
	viewerID string,
	chatID string,
) error {
	repo := repository.NewChatRepository(client)

	chat, err := repo.GetChat(viewerID, chatID)
	if err != nil {
		return err
	}

	websocketMsg := &models.WebSocketMessage{
		Type:     models.CHAT,
		Chat:     *chat,
		IsViewer: false,
	}

	bytes, err := json.Marshal(websocketMsg)
	if err != nil {
		return err
	}

	err = conn.Send(bytes)
	if err != nil {
		return err
	}

	return nil
}
