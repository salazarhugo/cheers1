package repository

import (
	"firebase.google.com/go/v4/messaging"
	"github.com/go-redis/redis/v9"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"os"
)

type Repository interface {
	SendNotification(
		userWithToken map[string][]string,
		data map[string]string,
	) error
	FollowUserNotification(
		user *user.User,
		followedUser *user.User,
	) error
	CreateRegistrationToken(userID string, token string) error
	GetUserTokens(userID string) ([]string, error)
	GetUsersTokens(userIDs []string) ([]string, error)
	RemoveExpiredTokens(userId string, tokens []string, responses []*messaging.SendResponse)
	DeleteToken(userId string, token string) error
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
