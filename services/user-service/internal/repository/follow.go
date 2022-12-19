package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
)

func (p *userRepository) FollowUser(
	userID string,
	otherUserID string,
) error {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/FollowUser.cql")
	if err != nil {
		return err
	}

	params := map[string]interface{}{
		"userID":      userID,
		"otherUserID": otherUserID,
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

	response, err = p.GetUser(userID, otherUserID)
	if err != nil {
		return err
	}
	otherUser := response.User

	err = pubsub.PublishProtoWithBinaryEncoding("user-topic", &user.UserEvent{
		Event: &user.UserEvent_Follow{
			&user.FollowUser{
				User:         currentUser,
				FollowedUser: otherUser,
			},
		},
	})

	if err != nil {
		return err
	}

	return nil
}
