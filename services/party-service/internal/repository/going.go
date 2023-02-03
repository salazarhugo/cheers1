package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
)

func (p *partyRepository) UpdateWatchStatus(
	userID string,
	partyID string,
	watchStatus party.WatchStatus,
) error {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/UpdateWatchStatus.cql")
	if err != nil {
		return err
	}

	params := map[string]interface{}{
		"userID":      userID,
		"partyID":     partyID,
		"watchStatus": watchStatus.String(),
	}

	_, err = session.Run(*cypher, params)
	if err != nil {
		return err
	}

	return nil
}
