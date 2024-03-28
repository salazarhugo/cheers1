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
	CreateMessage(
		msg *models.ChatMessage,
	) error

	ListChatMessage(
		chatID string,
		offset int,
		limit int,
	) ([]*models.ChatMessage, error)

	GetLastRead(
		chatID string,
		userID string,
	) (int64, error)

	SetLastRead(
		chatID string,
		userID string,
		timestamp int64,
	) error

	GetUnreadCount(
		chatID string,
		userID string,
	) (int64, error)

	ResetUnreadCount(
		roomID string,
		userID string,
	) error

	IncrementUnreadCount(
		roomID string,
		userID string,
	) error

	LikeMessage(chatID string, messageId string)
	UnlikeMessage(chatID string, messageId string)
	CreateGroup(name string, UUIDs []string) (*models.Chat, error)
	GetOrCreateDirectRoom(userId string, otherUserId string) (*models.Chat, error)
	LeaveChat(userId string, chatID string)
	IsMember(userId string, chatID string) (bool, error)
	IsAdmin(userId string, chatID string) (bool, error)
	DeleteRoom(chatID string) error
	GetChat(userID string, chatID string) (*models.Chat, error)
	ListRooms(userID string) ([]*models.Chat, error)
	ListChatIds(userID string) ([]string, error)
	GetRoomMembers(chatID string) ([]string, error)
	GetUserItem(userID string) (*user.UserItem, error)
	ListUser(userIDs []string) ([]*user.UserItem, error)
	DeleteTokens(userID string) int64
	AddToken(userID string, token string)
	GetTokens(userID string) []string
	GetOtherUserId(chatID string, userID string) (string, error)
	SetSeen(chatID string, userId string)
	GetRoomStatus(chatID string, userId string, otherUserId string) (pb.RoomStatus, error)
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

func (cache *redisCache) LikeMessage(chatID string, messageId string) {
	//TODO implement me
	panic("implement me")
}

func (cache *redisCache) UnlikeMessage(chatID string, messageId string) {
	//TODO implement me
	panic("implement me")
}
