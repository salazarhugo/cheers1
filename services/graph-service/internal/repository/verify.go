package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
)

func (p *userRepository) VerifyUser(
	userID string,
) error {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/VerifyUser.cql")
	if err != nil {
		return err
	}

	params := map[string]interface{}{
		"userID": userID,
	}

	_, err = session.Run(*cypher, params)
	if err != nil {
		return err
	}

	result, err := p.GetUser(userID, userID)

	err = pubsub.PublishProtoWithBinaryEncoding("user-topic", &pb.UserEvent{
		Event: &pb.UserEvent_Update{
			Update: &pb.UpdateUser{
				User: result.User,
			},
		},
	})

	return nil
}
