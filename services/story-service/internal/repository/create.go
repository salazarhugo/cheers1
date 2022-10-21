package repository

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	storypb "github.com/salazarhugo/cheers1/genproto/cheers/type/story"
	utils "github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

func (p *storyRepository) CreateStory(
	userID string,
	story *storypb.Story,
) (string, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/CreateStory.cql")
	if err != nil {
		return "", err
	}

	story.Id = uuid.NewString()
	story.CreateTime = timestamppb.Now()

	bytes, err := protojson.Marshal(story)
	if err != nil {
		return "", err
	}

	var m = make(map[string]interface{}, 0)

	err = json.Unmarshal(bytes, &m)
	if err != nil {
		return "", err
	}

	params := map[string]interface{}{
		"userID": userID,
		"story":  m,
	}

	_, err = session.Run(*cypher, params)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return story.Id, nil
}
