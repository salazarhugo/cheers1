package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	party "github.com/salazarhugo/cheers1/gen/go/cheers/type/party"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/mapper"
)

func (p *partyRepository) GetPartyWithSlug(slug string) (*party.Party, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/GetPartyWithSlug.cql")
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"slug": slug,
	}

	result, err := session.Run(*cypher, params)

	if err != nil {
		return nil, err
	}

	party := &party.Party{}

	if result.Next() {
		data := result.Record().Values[0]
		err := mapper.MapToProto(party, data)
		if err != nil {
			return nil, err
		}
	}

	return party, nil
}
