package cache

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
)

func (cache *redisCache) GetMessage(messageID string) ([]*pb.Message, error) {
	//res, err := cache.client.LIndex(context.Background(), getKeyUser(userId)).Result()
	return []*pb.Message{}, nil
}
