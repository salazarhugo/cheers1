package repository

import (
	"github.com/go-redis/redis/v9"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"os"
)

type Repository interface {
	CreateFriendRequest(userId string, friendId string) error
	ListFriendRequests(userId string) ([]*user.UserItem, error)
	ListFriend(userId string) ([]string, error)
	AcceptFriendRequest(userId string, friendId string) error
	DeleteFriendRequest(userId string, friendId string) error
	DeleteFriend(userId string, friendId string) error
	CheckFriend(user1 string, user2 string) (bool, error)
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
