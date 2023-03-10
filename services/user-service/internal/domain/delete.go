package domain

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/salazarhugo/cheers1/libs/utils"
	"log"
)

func (userService *UserService) DeleteUser(
	userID string,
) error {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/DeleteUser.cql")
	if err != nil {
		return err
	}

	params := map[string]interface{}{
		"userID": userID,
	}

	result, err := session.Run(*cypher, params)
	if err != nil {
		return err
	}
	log.Println(result)

	return nil
}
