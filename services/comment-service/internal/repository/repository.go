package repository

import (
	"github.com/go-redis/redis/v9"
	"github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"os"
)

type Repository interface {
	CreateComment(userId string, text string, postId string) (string, error)
	CreateReplyComment(userId string, text string, postId string, replyCommentId string) (string, error)
	UpdateUser(user *user.User) error
	GetLastComment(postId string) (*comment.CommentItem, error)
	GetComment(commentID string) (*comment.Comment, error)
	GetCommentItem(commentID string) (*comment.CommentItem, error)
	ListComment(postId string) ([]*comment.CommentItem, error)
	ListReplies(commentId string) ([]*comment.CommentItem, error)
	DeleteComment(commentId string) error

	GetUserItem(userId string) (*user.UserItem, error)
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
