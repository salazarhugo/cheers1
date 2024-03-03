package redisdb

import (
	"context"
	"github.com/go-redis/redis/v9"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/models"
	"time"
)

type RoomCache interface {
	SetMessage(msg *pb.Message) error
	ListMessage(
		roomId string,
		offset int,
		limit int,
	) []*pb.Message
	LikeMessage(roomId string, messageId string)
	UnlikeMessage(roomId string, messageId string)
	CreateGroup(name string, UUIDs []string) (*models.Chat, error)
	GetOrCreateDirectRoom(userId string, otherUserId string) (*models.Chat, error)
	LeaveRoom(userId string, roomId string)
	IsMember(userId string, roomId string) (bool, error)
	DeleteRoom(roomId string) error
	GetRoomWithId(userId string, roomId string) (*models.Chat, error)
	ListRooms(userID string) ([]*pb.Room, error)
	GetRoomMembers(roomId string) ([]string, error)
	GetUserItem(userId string) (*user.UserItem, error)
	ListUser(userIds []string) ([]*user.UserItem, error)
	DeleteTokens(userId string) int64
	AddToken(userId string, token string)
	GetTokens(userId string) []string
	GetOtherUserId(roomId string, userId string) (string, error)
	SetSeen(roomId string, userId string)
	GetRoomStatus(roomId string, userId string, otherUserId string) (pb.RoomStatus, error)
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
