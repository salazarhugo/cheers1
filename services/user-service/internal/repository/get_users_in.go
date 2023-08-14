package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/mapper"
)

func (p *userRepository) GetUsersIn(
	userIDs []string,
) ([]*user.User, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/GetUsersIn.cql")
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"userIDs": userIDs,
	}

	results, err := session.Run(*cypher, params)
	if err != nil {
		return nil, err
	}

	users := make([]*user.User, 0)

	for results.Next() {
		m := results.Record().Values[0]
		userItem := &user.User{}
		err := mapper.MapToProto(userItem, m)
		if err != nil {
			return nil, err
		}
		users = append(users, userItem)
	}

	return users, nil
}
