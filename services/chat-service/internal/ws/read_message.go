package ws

import (
	"encoding/json"
	"github.com/go-redis/redis/v9"
	"github.com/gorilla/websocket"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/models"
	"log"
)

func ReadMessage(
	conn *Connection,
	client *redis.Client,
	userID string,
) {
	for {
		messageType, b, err := conn.Socket.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		log.Printf("new message: %s, %d", string(b), messageType)

		var msg models.WebSocketMessage
		err = json.Unmarshal(b, &msg)
		if err != nil {
			log.Println(err)
			return
		}

		msg.UserId = userID

		HandleCommands(client, &msg)
	}
}
