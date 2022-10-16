package repository

import (
	"encoding/json"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/protobuf/encoding/protojson"
)

func (p *postRepository) UpdateUser(
	userID string,
	user *pb.UpdateUserRequest,
) error {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/UpdateUser.cql")
	if err != nil {
		return err
	}

	bytes, err := protojson.Marshal(user)
	if err != nil {
		return err
	}

	var m = make(map[string]interface{}, 0)

	err = json.Unmarshal(bytes, &m)
	if err != nil {
		return err
	}

	params := map[string]interface{}{
		"userID": userID,
		"user":   m,
	}

	_, err = session.Run(*cypher, params)
	if err != nil {
		return err
	}

	return nil
}
