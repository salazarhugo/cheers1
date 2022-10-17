package repository

import (
	"github.com/go-redis/redis/v9"
	"time"
)

type ChatRepository interface {
}

type partyRepository struct {
	expires time.Duration
	client  *redis.Client
}

func NewChatRepository(expires time.Duration, client *redis.Client) ChatRepository {
	return &partyRepository{
		expires: expires,
		client:  client,
	}
}
