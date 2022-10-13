package repository

import (
	"github.com/labstack/gommon/log"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	party "github.com/salazarhugo/cheers1/genproto/cheers/type/party"
	"github.com/salazarhugo/cheers1/libs/utils"
)

func (p *partyRepository) GetParty(id string) (*party.Party, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/CreateParty.cql")
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"userId":  "",
		"partyId": id,
	}

	result, err := session.Run(*cypher, params)

	if err != nil {
		return nil, err
	}

	if result.Next() {
		log.Info(result.Record().Values)
	}

	return &party.Party{}, nil
}
