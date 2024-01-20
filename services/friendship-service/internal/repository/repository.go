package repository

import (
	"github.com/go-redis/redis/v9"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"github.com/salazarhugo/cheers1/libs/utils"
	"gorm.io/gorm"
	"os"
)

type Repository interface {
	CreateFriend(
		userId string,
		friendId string,
	) error

	CreateFriendRequest(
		userId string,
		friendId string,
	) error

	GetFriendRequest(
		userID string,
		user2ID string,
	) (*FriendRequest, error)

	ListFriendRequests(
		page int,
		pageSize int,
		userId string,
	) ([]*user.UserItem, error)

	ListFriend(
		viewerID string,
		page int,
		pageSize int,
		userId string,
	) ([]*user.UserItem, error)

	AcceptFriendRequest(
		userId string,
		friendId string,
	) error

	DeleteFriendRequest(
		userId string,
		friendId string,
	) error

	DeleteFriend(
		userId string,
		friendId string,
	) error

	CheckFriend(
		user1 string,
		user2 string,
	) (bool, error)
}

type repository struct {
	spanner *gorm.DB
	redis   *redis.Client
}

func NewRepository() Repository {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ENDPOINT"),
		Password: os.Getenv("DB_PASSWORD"),
		DB:       0,
	})

	return &repository{
		spanner: utils.GetSpanner(),
		redis:   rdb,
	}
}
