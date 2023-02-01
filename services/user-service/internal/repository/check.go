package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/salazarhugo/cheers1/libs/utils"
)

func (p *userRepository) CheckUsername(
	username string,
) (bool, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/CheckUsername.cql")
	if err != nil {
		return false, err
	}

	params := map[string]interface{}{
		"username": username,
	}

	result, err := session.Run(*cypher, params)
	if err != nil {
		return false, err
	}

	if result.Next() {
		exists := result.Record().Values[0].(bool)
		return exists, nil
	}

	return false, nil
}
