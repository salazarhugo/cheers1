package repository

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/post/v1"
	"github.com/salazarhugo/cheers1/genproto/cheers/type/post"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/prototext"
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
		m := result.Record().Values[0]
		log.Info(m)
		bytes, err := json.Marshal(m)
		if err != nil {
			return nil, err
		}
		post := &pb.PostResponse{}
		log.Info(string(bytes))
		err = (prototext.UnmarshalOptions{DiscardUnknown: true}).Unmarshal(bytes, post)
		if err != nil {
			log.Info("Error unmarshal")
			return nil, err
		}
		posts = append(posts, post)
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
