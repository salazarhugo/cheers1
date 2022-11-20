package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
)

func (p *userRepository) GetUser(
	userID string,
	otherUserID string,
) (*pb.GetUserResponse, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/GetUser.cql")
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"userID":      userID,
		"otherUserID": otherUserID,
	}

	result, err := session.Run(*cypher, params)
	if err != nil {
		return nil, err
	}

	user := &pb.GetUserResponse{}

	if result.Next() {
		m := result.Record().Values[0]
		err := utils.MapToProto(user, m)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}
