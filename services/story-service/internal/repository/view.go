package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/story/v1"
	utils "github.com/salazarhugo/cheers1/libs/utils"
)

func (p *storyRepository) ViewStory(
	userID string,
	storyID string,
) (*pb.ViewStoryResponse, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/ViewStory.cql")
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

	return &pb.ViewStoryResponse{
		Success: true,
	}, nil
}
