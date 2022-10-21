package repository

import (
	"encoding/json"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/story/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/protobuf/encoding/protojson"
)

func (p *storyRepository) GetStory(
	userID string,
	storyID string,
) (*pb.StoryResponse, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/GetStory.cql")
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"userID":  userID,
		"storyID": storyID,
	}

	result, err := session.Run(*cypher, params)
	if err != nil {
		return nil, err
	}

	story := &pb.StoryResponse{}

	if result.Next() {
		m := result.Record().Values[0]
		bytes, err := json.Marshal(m)
		if err != nil {
			return nil, err
		}
		err = protojson.Unmarshal(bytes, story)
		if err != nil {
			return nil, err
		}
	}

	return story, nil
}
