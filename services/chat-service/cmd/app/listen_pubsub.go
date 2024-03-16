package main

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v9"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/models"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/repository"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/ws"
	"log"
)

func ListenMessages(
	conn *ws.Connection,
	client *redis.Client,
	viewerID string,
) {
	repo := repository.NewChatRepository(client)
	rooms, _ := repo.GetInbox(viewerID)

	for _, room := range rooms {
		roomId := room.Room.Id
		go ListenRoom(conn, client, roomId, viewerID)
	}
}

func ListenRoom(
	conn *ws.Connection,
	client *redis.Client,
	roomId string,
	viewerID string,
) {
	sub := client.Subscribe(context.Background(), roomId)
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

//var chatMessage chat.Message
//err := protojson.Unmarshal([]byte(msg.Payload), &chatMessage)
//chatMessage.Status = chat.Message_DELIVERED
//bytes, err := protojson.Marshal(&chatMessage)
