package repository

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/protobuf/encoding/protojson"
)

func (p *userRepository) ListFollowers(
	userID string,
	request *pb.ListFollowersRequest,
) ([]*user.UserItem, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/ListFollowers.cql")
	if err != nil {
		return nil, err
	}

	skip := request.GetPageSize() * request.GetPage()

	params := map[string]interface{}{
		"userID":      userID,
		"otherUserID": request.GetUserId(),
		"skip":        skip,
		"pageSize":    request.GetPageSize(),
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
