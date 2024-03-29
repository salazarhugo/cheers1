package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/salazarhugo/cheers1/libs/utils"
)

func (p *userRepository) UnfollowUser(
	userID string,
	otherUser string,
) error {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/UnfollowUser.cql")
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

	return nil
}
