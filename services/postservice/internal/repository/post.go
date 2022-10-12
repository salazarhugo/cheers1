package repository

import (
	"github.com/labstack/gommon/log"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/post/v1"
	"github.com/salazarhugo/cheers1/genproto/cheers/type/post"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
)

type PostRepository interface {
	CreatePost(post postpb.Post) error
	GetPost(id string) (*postpb.Post, error)
	ListPost(userID string, request *pb.ListPostRequest) (*pb.ListPostResponse, error)
	UpdatePost(post postpb.Post) error
	DeletePost(id string) error
}

type postRepository struct {
	driver neo4j.Driver
}

func NewPostRepository(driver neo4j.Driver) PostRepository {
	return &postRepository{driver: driver}
}

func (p *postRepository) CreatePost(post postpb.Post) error {
	return nil
}

func (p *postRepository) GetPost(id string) (*postpb.Post, error) {
	return &postpb.Post{}, nil
}

func (p *postRepository) UpdatePost(post postpb.Post) error {
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

	bytes, err := os.ReadFile("internal/cql/ListPost.cql")
	if err != nil {
		return nil, status.Error(codes.Internal, "failed reading cql file")
	}
	cypher := string(bytes)

	params := map[string]interface{}{
		"userId":   userID,
		"skip":     0,
		"pageSize": pageSize,
	}

	result, err := session.Run(cypher, params)
	if err != nil {
		return nil, err
	}

	// Empty list
	posts := make([]*pb.PostResponse, 0)

	for result.Next() {
		log.Info(result.Record())
		//posts = append(posts, result.Record().Values[0].(*pb.PostResponse))
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
