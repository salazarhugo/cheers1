package ws

import (
	"github.com/go-redis/redis/v9"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/repository"
)

func SubscribePubSub(
	conn *Connection,
	client *redis.Client,
	viewerID string,
) {
	repo := repository.NewChatRepository(client)
	chatIDs, _ := repo.ListChatsIds(viewerID)

	go SubscribeUser(conn, client, viewerID)

	for _, chatID := range chatIDs {
		go SubscribeChat(conn, client, chatID, viewerID)
	}
}

//var chatMessage chat.Message
//err := protojson.Unmarshal([]byte(msg.Payload), &chatMessage)
//chatMessage.Status = chat.Message_DELIVERED
//bytes, err := protojson.Marshal(&chatMessage)
