package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	userpb "github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"log"
	"time"
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
	user.CreateTime = time.Now().Unix()

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

	err = pubsub.PublishProtoWithBinaryEncoding("user-topic", &userpb.UserEvent{
		Event: &userpb.UserEvent_Create{
			Create: &userpb.CreateUser{
				User: user,
			},
		},
	})

	if err != nil {
		return "", err
	}

	return user.Id, nil
}
