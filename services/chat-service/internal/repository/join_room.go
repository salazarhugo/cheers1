package repository

import (
	"context"
	json2 "encoding/json"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"log"
)

func (c chatRepository) JoinRoom(
	request *pb.JoinRoomRequest,
	server pb.ChatService_JoinRoomServer,
) error {
	messages := c.cache.ListMessage(request.GetRoomId(), 10)
	log.Println(messages)
	for _, msg := range messages {
		if err := server.Send(msg); err != nil {
			log.Printf("send error %v", err)
		}
	}

	sub := c.cache.Subscribe(context.Background(), request.GetRoomId())
	defer sub.Close()

	ch := sub.Channel()

	for {
		select {
		case <-server.Context().Done():
			return nil
		case msg := <-ch:
			message := pb.Message{}
			if err := json2.Unmarshal([]byte(msg.Payload), &message); err != nil {
				log.Println(err)
				return err
			}

			log.Println("sending message to stream: " + message.Text)
			err := server.Send(&message)
			if err != nil {
				log.Println(err)
				return err
			}
		}
	}

	return nil
}
