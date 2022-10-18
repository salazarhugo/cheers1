package repository

import (
	"context"
	json2 "encoding/json"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/chat/v1"
	"log"
)

func (c chatRepository) JoinRoom(request *pb.JoinRoomRequest, server pb.ChatService_JoinRoomServer) error {
	ctx := context.Background()

	messages := c.cache.GetMessages(request.GetRoomId())
	for _, msg := range messages {
		if err := server.Send(msg); err != nil {
			log.Printf("send error %v", err)
		}
	}

	sub := c.cache.Subscribe(ctx, request.GetRoomId())
	defer sub.Close()

	for {
		msg, err := sub.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
		message := pb.Message{}
		if err := json2.Unmarshal([]byte(msg.Payload), &message); err != nil {
			panic(err)
		}
		log.Println(message)
		err = server.Send(&message)
		if err != nil {
			return err
		}
	}
}
