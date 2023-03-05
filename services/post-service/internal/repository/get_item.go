package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
)

func (p *postRepository) GetPostItem(
	userID string,
	postID string,
) (*pb.PostResponse, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/GetPostItem.cql")
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"userID": userID,
		"postID": postID,
	}

	result, err := session.Run(*cypher, params)
	if err != nil {
		return nil, err
	}

	post := &pb.PostResponse{}

	if result.Next() {
		m := result.Record().Values[0]
		err := utils.MapToProto(post, m)
		if err != nil {
			return nil, err
		}
	}

	return post, nil
}
