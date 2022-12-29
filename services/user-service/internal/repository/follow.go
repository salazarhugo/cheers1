package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
)

func (p *userRepository) FollowUser(
	userID string,
	otherUser string,
) error {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/FollowUser.cql")
	if err != nil {
		return err
	}

	params := map[string]interface{}{
		"userID":    userID,
		"otherUser": otherUser,
	}

	_, err = session.Run(*cypher, params)
	if err != nil {
		return err
	}
	response, err := p.GetUser(userID, userID)
	if err != nil {
		return err
	}
	currentUser := response.User

	response, err = p.GetUser(userID, otherUser)
	if err != nil {
		return err
	}
	followedUser := response.User

	err = pubsub.PublishProtoWithBinaryEncoding("user-topic", &user.UserEvent{
		Event: &user.UserEvent_Follow{
			Follow: &user.FollowUser{
				User:         currentUser,
				FollowedUser: followedUser,
			},
		},
	})

	if err != nil {
		return err
	}

	return nil
}
