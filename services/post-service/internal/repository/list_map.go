package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	utils "github.com/salazarhugo/cheers1/libs/utils"
)

func (p *postRepository) ListMapPost(
	userID string,
	request *pb.ListMapPostRequest,
) (*pb.ListMapPostResponse, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	pageSize := request.GetPageSize()
	page := request.GetPage()

	skip := page * pageSize

	cypher, err := utils.GetCypher("internal/queries/ListMapPost.cql")
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"userID":           userID,
		"userIdOrUsername": userID,
		"skip":             int(skip),
		"pageSize":         int(pageSize),
	}

	result, err := session.Run(*cypher, params)
	if err != nil {
		return nil, err
	}

	// Empty ListMap
	posts := make([]*pb.PostResponse, 0)

	for result.Next() {
		m := result.Record().Values[0]
		post := &pb.PostResponse{}
		err := utils.MapToProto(post, m)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return &pb.ListMapPostResponse{
		Posts: posts,
	}, nil
}
