package repository

import (
	"github.com/go-redis/redis/v9"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	postpb "github.com/salazarhugo/cheers1/gen/go/cheers/type/post"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"github.com/salazarhugo/cheers1/libs/utils"
	"gorm.io/gorm"
	"os"
)

type PostRepository interface {
	CreatePost(userID string, post *Post) (string, error)
	GetPost(postID string) (*postpb.Post, error)
	GetPostById(postID string) (*Post, error)
	GetPostItem(userID string, postID string) (*pb.PostResponse, error)
	UpdatePostLastComment(user *user.UserItem, comment *comment.Comment) error
	DeletePost(userID string, postID string) error
	DeletePostById(postID string) error

	FeedPost(userID string, request *pb.FeedPostRequest) (*pb.FeedPostResponse, error)
	ListPost(userID string, request *pb.ListPostRequest) (*pb.ListPostResponse, error)
	ListMapPost(userID string, request *pb.ListMapPostRequest) (*pb.ListMapPostResponse, error)
	LikePost(userID string, postID string) (*pb.LikePostResponse, error)
	UnlikePost(userID string, postID string) (*pb.UnlikePostResponse, error)
	IncrementCommentCount(postID string) error
	DecrementCommentCount(postID string) error
	GetCommentCount(postID string) (int, error)
}

type postRepository struct {
	driver  neo4j.Driver
	redis   *redis.Client
	spanner *gorm.DB
}

func NewPostRepository() PostRepository {
	driver := utils.GetDriver()
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ENDPOINT"),
		Password: os.Getenv("DB_PASSWORD"),
		DB:       0,
	})

	return &postRepository{
		driver:  driver,
		redis:   rdb,
		spanner: utils.GetSpanner(),
	}
}
