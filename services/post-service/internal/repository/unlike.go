package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	utils "github.com/salazarhugo/cheers1/libs/utils"
)

func (p *postRepository) UnlikePost(
	userID string,
	postID string,
) (*pb.UnlikePostResponse, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/UnlikePost.cql")
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"userID": userID,
		"postID": postID,
	}

	_, err = session.Run(*cypher, params)
	if err != nil {
		return nil, err
	}

	return &pb.UnlikePostResponse{
		Success: true,
	}, nil
}
