package repository

import (
	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	party "github.com/salazarhugo/cheers1/gen/go/cheers/type/party"
	"github.com/salazarhugo/cheers1/libs/utils"
	"log"
	"time"
)

func (p *partyRepository) CreateParty(
	userID string,
	party *party.Party,
) (*party.Party, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/CreateParty.cql")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	party.Id = uuid.NewString()
	party.CreateTime = time.Now().Unix()

	m, err := utils.ProtoToMap(party)
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"userID": userID,
		"party":  m,
	}

	_, err = session.Run(*cypher, params)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return party, nil
}
