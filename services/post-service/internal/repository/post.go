package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/post/v1"
	"github.com/salazarhugo/cheers1/genproto/cheers/type/post"
)

type PostRepository interface {
	CreatePost(userID string, post *postpb.Post) (string, error)
	GetPost(userID string, postID string) (*pb.PostResponse, error)
	UpdatePost(post *postpb.Post) (*pb.PostResponse, error)
	DeletePost(id string) error

	FeedPost(userID string, request *pb.FeedPostRequest) (*pb.FeedPostResponse, error)
	LikePost(userID string, postID string) (*pb.LikePostResponse, error)
	UnlikePost(userID string, postID string) (*pb.UnlikePostResponse, error)
}

type postRepository struct {
	driver neo4j.Driver
}

func (p *postRepository) UpdatePost(post *postpb.Post) (*pb.PostResponse, error) {
	panic("implement me")
}

func NewPostRepository(driver neo4j.Driver) PostRepository {
	return &postRepository{driver: driver}
}

func (p *postRepository) DeletePost(id string) error {
	return nil
}
