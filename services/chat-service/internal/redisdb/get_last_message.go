package redisdb

import (
	"context"
	json2 "encoding/json"
	"github.com/go-redis/redis/v9"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"strings"
)

func getLatestMessage(
	client *redis.Client,
	roomId string,
) (*pb.Message, error) {
	lastMessage, err := client.ZRevRangeByScore(
		context.Background(),
		getKeyRoomMessages(roomId),
		&redis.ZRangeBy{
			Min:    "-inf",
			Max:    "+inf",
			Offset: 0,
			Count:  1,
		}).Result()

	if len(lastMessage) == 0 {
		return nil, nil
	}

	message := &pb.Message{}
	dec := json2.NewDecoder(strings.NewReader(lastMessage[0]))
	err = dec.Decode(message)
	if err != nil {
		return nil, err
	}

	return message, err
}
