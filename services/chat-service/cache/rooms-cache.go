package cache

import (
	"context"
	"github.com/go-redis/redis/v9"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/chat/v1"
	"time"
)

type RoomCache interface {
	SetMessage(msg *pb.Message)
	GetMessages(roomId string) []*pb.Message
	LikeMessage(roomId string, messageId string)
	UnlikeMessage(roomId string, messageId string)
	CreateGroup(name string, UUIDs []string) *pb.Room
	GetOrCreateDirectRoom(userId string, otherUserId string) *pb.Room
	LeaveRoom(userId string, roomId string)
	IsMember(userId string, roomId string) bool
	DeleteRoom(roomId string)
	GetRoomWithId(userId string, roomId string) *pb.Room
	GetRooms(userId string) []*pb.Room
	GetRoomMembers(roomId string) []string
	GetUser(userId string) (interface{}, error)
	DeleteTokens(userId string) int64
	AddToken(userId string, token string)
	GetTokens(userId string) []string
	GetOtherUserId(roomId string, userId string) (string, error)
	SetSeen(roomId string, userId string)
	GetRoomStatus(roomId string, userId string, otherUserId string) pb.RoomStatus
	Subscribe(ctx context.Context, channel string) *redis.PubSub
	Publish(ctx context.Context, channel string, message interface{}) *redis.IntCmd
}

type redisCache struct {
	expires time.Duration
	client  *redis.Client
}

func NewCache(expires time.Duration, client *redis.Client) RoomCache {
	return &redisCache{
		expires: expires,
		client:  client,
	}
}

func (cache *redisCache) Publish(ctx context.Context, channel string, message interface{}) *redis.IntCmd {
	return cache.client.Publish(ctx, channel, message)
}

func (cache *redisCache) Subscribe(ctx context.Context, channel string) *redis.PubSub {
	return cache.client.Subscribe(ctx, channel)
}

func (cache *redisCache) LikeMessage(roomId string, messageId string) {
	//TODO implement me
	panic("implement me")
}

func (cache *redisCache) UnlikeMessage(roomId string, messageId string) {
	//TODO implement me
	panic("implement me")
}
