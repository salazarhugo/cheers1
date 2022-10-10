package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/salazarhugo/cheers1/genproto/cheers/type/post"
	pb "github.com/salazarhugo/cheers1/services/postservice/genproto/cheers/post/v1"
)

type PostRepository interface {
	CreatePost(post post.Post) error
	GetPost(id string) (*post.Post, error)
	ListPost(userID string, pageToken string) (*pb.ListPostResponse, error)
	UpdatePost(post post.Post) error
	DeletePost(id string) error
}

type postRepository struct {
	session neo4j.Session
}

func NewPostRepository(session neo4j.Session) PostRepository {
	return &postRepository{session: session}
}

func (p *postRepository) CreatePost(post post.Post) error {
	return nil
}

func (p *postRepository) GetPost(id string) (*post.Post, error) {
	return &post.Post{}, nil
}

func (p *postRepository) UpdatePost(post post.Post) error {
	return nil
}

func (p *postRepository) DeletePost(id string) error {
	return nil
}

func (p *postRepository) ListPost(userID string, pageToken string) (*pb.ListPostResponse, error) {
	return &pb.ListPostResponse{
		Posts:         nil,
		NextPageToken: "",
	}, nil
}
