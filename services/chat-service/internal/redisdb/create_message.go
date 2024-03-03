package redisdb

import (
	"bytes"
	"context"
	json2 "encoding/json"
	"github.com/go-redis/redis/v9"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"time"
)

func (cache *redisCache) SetMessage(
	msg *pb.Message,
) error {
	client := cache.client
	ctx := context.Background()
	buff := bytes.NewBufferString("")
	enc := json2.NewEncoder(buff)
	err := enc.Encode(msg)
	if err != nil {
		return err
	}

	err = client.Watch(ctx, func(tx *redis.Tx) error {
		cmd := tx.ZAdd(ctx, getKeyRoomMessages(msg.GetRoomId()), redis.Z{
			Score:  float64(msg.CreateTime),
			Member: buff.String(),
		})

		if cmd.Err() != nil {
			return cmd.Err()
		}

		cmd2 := tx.Expire(context.Background(), getKeyRoomMessages(msg.GetRoomId()), 24*time.Hour)
		if cmd2.Err() != nil {
			return cmd.Err()
		}

		return nil
	})

	return err
}
