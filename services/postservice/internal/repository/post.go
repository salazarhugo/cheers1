package repository

import (
	"encoding/json"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/post/v1"
	"github.com/salazarhugo/cheers1/genproto/cheers/type/post"
	utils "github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
)

type PostRepository interface {
	CreatePost(userID string, post *postpb.Post) error
	GetPost(id string) (*postpb.Post, error)
	ListPost(userID string, request *pb.ListPostRequest) (*pb.ListPostResponse, error)
	UpdatePost(post *postpb.Post) error
	DeletePost(id string) error
}

type postRepository struct {
	driver neo4j.Driver
}

func NewPostRepository(driver neo4j.Driver) PostRepository {
	return &postRepository{driver: driver}
}

func (p *postRepository) CreatePost(userID string, post *postpb.Post) error {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/cql/CreatePost.cql")
	if err != nil {
		return err
	}

	bytes, err := protojson.Marshal(post)
	if err != nil {
		return err
	}

	var m = make(map[string]interface{}, 0)

	err = json.Unmarshal(bytes, &m)
	if err != nil {
		return err
	}

	params := map[string]interface{}{
		"userID": userID,
		"post":   m,
	}

	_, err = session.Run(*cypher, params)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (p *postRepository) GetPost(id string) (*postpb.Post, error) {
	return &postpb.Post{}, nil
}

func (p *postRepository) UpdatePost(post *postpb.Post) error {
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

	cypher, err := GetCypher("internal/cql/ListPost.cql")
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"userID":   userID,
		"skip":     0,
		"pageSize": pageSize,
	}

	result, err := session.Run(*cypher, params)
	if err != nil {
		return nil, err
	}

	// Empty list
	posts := make([]*pb.PostResponse, 0)

	for result.Next() {
		m := result.Record().Values[0]
		bytes, err := json.Marshal(m)
		if err != nil {
			return nil, err
		}
		post := &pb.PostResponse{}
		err = protojson.Unmarshal(bytes, post)
		if err != nil {
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
