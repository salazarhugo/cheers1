package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
)

func (p *partyRepository) GetPartyItem(
	userID string,
	partyID string,
) (*pb.PartyItem, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/GetPartyItem.cql")
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"userID":  userID,
		"partyID": partyID,
	}

	result, err := session.Run(*cypher, params)

	if err != nil {
		return nil, err
	}

	item := &pb.PartyItem{}

	if result.Next() {
		data := result.Record().Values[0]
		err := utils.MapToProto(item, data)
		if err != nil {
			return nil, err
		}
	}

	return item, nil
}
