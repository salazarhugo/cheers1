package repository

import (
	"context"
	"encoding/json"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"log"
)

func (c chatRepository) SendMessage(
	msg *pb.Message,
) error {
	ctx := context.Background()

	// Store message in cache
	err := c.cache.SetMessage(msg)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(msg)
	// Redis Pub/Sub
	c.cache.Publish(ctx, msg.RoomId, bytes)
	// Google Cloud Pub/Sub
	go PublishToPubSub(msg.SenderId, msg.RoomId)

	return nil
}

func PublishToPubSub(senderID string, roomID string) {
	err := utils.PublishProtoMessages("chat-topic", &pb.ChatEvent{
		SenderId: senderID,
		RoomId:   roomID,
	})
	if err != nil {
		log.Println(err)
	}
}
