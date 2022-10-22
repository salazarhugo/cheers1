package repository

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/salazarhugo/cheers1/genproto/cheers/type/user"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/protobuf/encoding/protojson"
)

func (p *postRepository) GetUsersIn(
	userID string,
	userIDs []string,
) ([]*user.UserItem, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/GetUsersIn.cql")
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"userID":  userID,
		"userIDs": userIDs,
	}

	results, err := session.Run(*cypher, params)
	if err != nil {
		return nil, err
	}

	users := make([]*user.UserItem, 0)

	for results.Next() {
		m := results.Record().Values[0]
		bytes, err := json.Marshal(m)
		if err != nil {
			return nil, err
		}
		userItem := &user.UserItem{}
		err = protojson.Unmarshal(bytes, userItem)
		if err != nil {
			log.Error(err)
		}
		users = append(users, userItem)
	}

	return users, nil
}