package repository

import (
	"context"
	"github.com/google/uuid"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"time"
)

func (c chatRepository) SendMessage(
	msg *pb.Message,
) (*pb.Message, error) {
	ctx := context.Background()
	msg.Id = uuid.NewString()
	now := time.Now()
	msg.CreateTime = now.Unix()

	// Store message in cache
	err := c.cache.SetMessage(msg)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	bytes, err := protojson.Marshal(msg)
	// Redis Pub/Sub
	c.cache.Publish(ctx, msg.RoomId, bytes)
	// Google Cloud Pub/Sub
	//go PublishToPubSub(msg.SenderId, msg.RoomId)

	return msg, nil
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
