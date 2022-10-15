package repository

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/salazarhugo/cheers1/genproto/cheers/type/post"
	utils "github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

func (p *postRepository) CreatePost(
	userID string,
	post *postpb.Post,
) (string, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/CreatePost.cql")
	if err != nil {
		return "", err
	}

	post.Id = uuid.NewString()
	post.CreateTime = timestamppb.Now()

	bytes, err := protojson.Marshal(post)
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
		"post":   m,
	}

	_, err = session.Run(*cypher, params)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return post.Id, nil
}
