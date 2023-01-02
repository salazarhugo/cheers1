package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	postpb "github.com/salazarhugo/cheers1/gen/go/cheers/type/post"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"github.com/salazarhugo/cheers1/libs/utils"
)

type PostRepository interface {
	CreatePost(userID string, post *postpb.Post) (string, error)
	GetPost(userID string, postID string) (*pb.PostResponse, error)
	UpdatePostLastComment(user *user.UserItem, comment *comment.Comment) error
	DeletePost(userID string, postID string) error

	FeedPost(userID string, request *pb.FeedPostRequest) (*pb.FeedPostResponse, error)
	ListPost(userID string, request *pb.ListPostRequest) (*pb.ListPostResponse, error)
	ListMapPost(userID string, request *pb.ListMapPostRequest) (*pb.ListMapPostResponse, error)
	LikePost(userID string, postID string) (*pb.LikePostResponse, error)
	UnlikePost(userID string, postID string) (*pb.UnlikePostResponse, error)
}

type postRepository struct {
	driver neo4j.Driver
}

func NewPostRepository() PostRepository {
	return &postRepository{driver: utils.GetDriver()}
}
