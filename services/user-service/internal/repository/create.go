package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/salazarhugo/cheers1/genproto/cheers/type/user"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

func (p *userRepository) CreateUser(
	userID string,
	user *user.User,
) (string, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/CreateUser.cql")
	if err != nil {
		return "", err
	}

	user.Id = userID
	user.CreateTime = timestamppb.Now()

	m, err := utils.ProtoToMap(user)
	if err != nil {
		log.Println(err)
		return "", err
	}

	params := map[string]interface{}{
		"user": m,
	}

	_, err = session.Run(*cypher, params)

	if err != nil {
		log.Println(err)
		return "", err
	}

	return user.Id, nil
}
