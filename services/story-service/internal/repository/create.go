package repository

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/story/v1"
	utils "github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"time"
)

func (p *storyRepository) CreateStory(
	userID string,
	story *pb.Story,
) (string, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/CreateStory.cql")
	if err != nil {
		return "", err
	}

	story.Id = uuid.NewString()
	story.CreateTime = time.Now().Unix()

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
