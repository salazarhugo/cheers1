package repository

import (
	"github.com/go-redis/redis/v9"
	"os"
)

type Repository interface {
	SendChatNotification(userWithToken map[string][]string) error
	CreateRegistrationToken(userID string, token string) error
	GetUserTokens(userID string) ([]string, error)
}

type repository struct {
	redis *redis.Client
}

func NewRepository() Repository {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ENDPOINT"),
		Password: os.Getenv("DB_PASSWORD"),
		DB:       0,
	})

	return &repository{
		redis: rdb,
	}
}
