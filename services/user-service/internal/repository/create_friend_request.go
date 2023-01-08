package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/salazarhugo/cheers1/libs/utils"
	"log"
)

func (p *userRepository) CreateFriendRequest(from string, to string) error {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/FriendRequestUser.cql")
	if err != nil {
		return err
	}

	params := map[string]interface{}{
		"userID":      from,
		"otherUserID": to,
	}

	_, err = session.Run(*cypher, params)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
