package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/salazarhugo/cheers1/genproto/cheers/type/post"
	pb "github.com/salazarhugo/cheers1/services/postservice/genproto/cheers/post/v1"
	"os"
)

type PostRepository interface {
	CreatePost(post post.Post) error
	GetPost(id string) (*post.Post, error)
	ListPost(userID string, request *pb.ListPostRequest) (*pb.ListPostResponse, error)
	UpdatePost(post post.Post) error
	DeletePost(id string) error
}

type postRepository struct {
	driver neo4j.Driver
}

func NewPostRepository(driver neo4j.Driver) PostRepository {
	return &postRepository{driver: driver}
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

func (p *postRepository) ListPost(
	userID string,
	request *pb.ListPostRequest,
) (*pb.ListPostResponse, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	pageSize := request.GetPageSize()
	//page := request.GetPageToken()
	//skip := page * pageSize

	bytes, err := os.ReadFile("../cql/ListPost.cql")
	if err != nil {
		return nil, err
	}
	cypher := string(bytes)

	params := map[string]interface{}{
		"userId":   userID,
		"skip":     0,
		"pageSize": pageSize,
	}

	result, err := session.Run(cypher, params)

	posts := make([]*pb.PostResponse, 0)

	for result.Next() {
		posts = append(posts, result.Record().Values[0].(*pb.PostResponse))
	}

	nextPageToken := ""
	if len(posts) > 0 {
		nextPageToken = posts[len(posts)-1].Post.Id
	}

	return &pb.ListPostResponse{
		Posts:         posts,
		NextPageToken: nextPageToken,
	}, nil
}
