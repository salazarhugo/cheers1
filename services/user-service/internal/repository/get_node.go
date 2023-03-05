package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (p *userRepository) GetUserNode(
	userId string,
) (*user.User, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/GetUserNode.cql")
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"userID": userId,
	}

	result, err := session.Run(*cypher, params)
	if err != nil {
		return nil, err
	}

	user := &user.User{}

	if result.Next() {
		m := result.Record().Values[0]
		err := utils.MapToProto(user, m)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, status.Error(codes.NotFound, "User not found")
	}

	return user, nil
}
