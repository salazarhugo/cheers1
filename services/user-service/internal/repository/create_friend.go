package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	userpb "github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"log"
)

func (p *userRepository) CreateFriend(from string, to string) error {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/FriendUser.cql")
	if err != nil {
		return err
	}

	params := map[string]interface{}{
		"userID":      from,
		"otherUserID": to,
	}

	_, err = session.Run(*cypher, params)
	if err != nil {
		return err
	}

	users, err := p.GetUsersIn([]string{from, to})
	if err != nil {
		log.Println(err)
		return err
	}

	err = pubsub.PublishProtoWithBinaryEncoding("user-topic", &userpb.UserEvent{
		Event: &userpb.UserEvent_Follow{
			Follow: &userpb.FollowUser{
				User:         users[0],
				FollowedUser: users[1],
			},
		},
	})

	return nil
}
