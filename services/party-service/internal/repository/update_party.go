package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	party "github.com/salazarhugo/cheers1/gen/go/cheers/type/party"
	"github.com/salazarhugo/cheers1/libs/utils"
)

func (p *partyRepository) UpdateParty(
	party *party.Party,
) (string, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/UpdateParty.cql")
	if err != nil {
		return "", err
	}

	m, err := utils.ProtoToMap(party)
	if err != nil {
		return "", err
	}

	params := map[string]interface{}{
		"partyID": party.Id,
		"party":   m,
	}

	_, err = session.Run(*cypher, params)
	if err != nil {
		return "", err
	}

	return party.Id, nil
}
