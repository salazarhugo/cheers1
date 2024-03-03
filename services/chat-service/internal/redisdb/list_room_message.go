package redisdb

import (
	"context"
	json2 "encoding/json"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"strings"
)

func (cache *redisCache) ListMessage(
	roomId string,
	offset int,
	limit int,
) []*pb.Message {
	client := cache.client

	values, err := client.ZRevRange(
		context.Background(),
		getKeyRoomMessages(roomId),
		int64(offset),
		int64(offset+limit),
	).Result()
	if err != nil {
		panic(err)
	}
	var messages []*pb.Message

	for i := range values {
		message := &pb.Message{}
		dec := json2.NewDecoder(strings.NewReader(values[i]))
		err := dec.Decode(message)
		if err != nil {
			panic(err)
		}
		item, _ := cache.GetUserItem(message.GetSenderId())
		message.SenderName = item.GetName()
		message.SenderPicture = item.GetPicture()
		message.SenderUsername = item.GetUsername()
		message.Status = pb.Message_DELIVERED
		messages = append(messages, message)
	}

	return messages
}
