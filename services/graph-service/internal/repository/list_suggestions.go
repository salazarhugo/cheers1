package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"github.com/salazarhugo/cheers1/libs/utils"
	"log"
)

func (p *userRepository) ListSuggestions(
	userID string,
) ([]*user.UserItem, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/ListSuggestions.cql")
	if err != nil {
		return nil, err
	}

	skip := 0
	pageSize := 10

	params := map[string]interface{}{
		"userID":   userID,
		"skip":     skip,
		"pageSize": pageSize,
	}

	results, err := session.Run(*cypher, params)
	if err != nil {
		return nil, err
	}

	userIDs := make([]string, 0)

	for results.Next() {
		m := results.Record().Values[0]
		userIDs = append(userIDs, m.(string))
	}

	log.Println(userIDs)
	items, err := p.GetUserItemsIn(userID, userIDs)
	if err != nil {
		return nil, err
	}

	return items, nil
}
