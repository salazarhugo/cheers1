package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/story/v1"
	utils "github.com/salazarhugo/cheers1/libs/utils"
)

func (p *storyRepository) LikeStory(
	userID string,
	storyID string,
) (*pb.LikeStoryResponse, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/LikeStory.cql")
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"userID":  userID,
		"storyID": storyID,
	}

	_, err = session.Run(*cypher, params)
	if err != nil {
		return nil, err
	}

	return &pb.LikeStoryResponse{
		Success: true,
	}, nil
}
