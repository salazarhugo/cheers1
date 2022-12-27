package repository

import (
	"github.com/go-redis/redis/v9"
	"github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"os"
)

type Repository interface {
	CreateComment(text string, postId string) error
	UpdateUser(user *user.User) error
	ListComment(postId string) ([]*comment.Comment, error)
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
