package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/story/v1"
	utils "github.com/salazarhugo/cheers1/libs/utils"
)

func (p *storyRepository) UnlikeStory(
	userID string,
	storyID string,
) (*pb.UnlikeStoryResponse, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/UnlikeStory.cql")
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

	return &pb.UnlikeStoryResponse{
		Success: true,
	}, nil
}
