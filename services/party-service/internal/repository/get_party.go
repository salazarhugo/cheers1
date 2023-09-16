package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	party "github.com/salazarhugo/cheers1/gen/go/cheers/type/party"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/mapper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (p *partyRepository) GetParty(id string) (*party.Party, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/GetParty.cql")
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"partyId": id,
	}

	result, err := session.Run(*cypher, params)

	if err != nil {
		return nil, err
	}

	if result.Next() {
		party := &party.Party{}
		data := result.Record().Values[0]
		err := mapper.MapToProto(party, data)
		if err != nil {
			return nil, err
		}
		return party, nil
	} else {
		return nil, status.Error(codes.NotFound, "party not found")
	}
}
