package ws

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v9"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/models"
	"log"
)

func HandleCommands(
	client *redis.Client,
	msg *models.WebSocketMessage,
) {
	ctx := context.Background()

	bytes, err := json.Marshal(&msg)
	if err != nil {
		log.Println(err)
	}

	switch msg.Type {
	case models.MESSAGE:
	case models.TYPING:
		//isTyping := msg.Typing.IsTyping
		err := client.Publish(ctx, msg.Typing.ChatId, string(bytes)).Err()
		if err != nil {
			return
		}
	case models.PRESENCE:
		//isPresent := msg.Presence.IsPresent
		err := client.Publish(ctx, msg.Presence.ChatId, string(bytes)).Err()
		if err != nil {
			return
		}
	}
	//var chatMessage chat.Message
	//err = proto.Unmarshal(b, &chatMessage)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//
	//msg, err := repo.SendMessage(chatMessage.SenderId, chatMessage.RoomId, chatMessage.Text)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//// Acknowledge message
	//msg.Status = chat.Message_SENT
	//bytes, err := protojson.Marshal(msg)
	//err = conn.Send(bytes)
	//if err != nil {
	//	log.Println(err)
	//}
}
