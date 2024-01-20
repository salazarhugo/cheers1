package repository

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/salazarhugo/cheers1/gen/go/cheers/location/v1"
	"os"
)

type Repository interface {
	UpdateLocation(
		userId string,
		latitude float64,
		longitude float64,
	) error

	UpdateGhostMode(
		userId string,
		ghostMode bool,
	) error

	GetGhostMode(
		userId string,
	) (bool, error)

	ListFriendLocation(
		ctx context.Context,
		userId string,
	) ([]*location.Location, error)
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
	pong, err := rdb.Ping(context.Background()).Result()
	fmt.Println(pong, err)

	return &repository{
		redis: rdb,
	}
}
