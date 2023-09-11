package repository

import (
	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	party "github.com/salazarhugo/cheers1/gen/go/cheers/type/party"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/mapper"
	"log"
	"time"
)

func (p *partyRepository) InsertParty(
	party *party.Party,
) (string, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/InsertParty.cql")
	if err != nil {
		log.Println(err)
		return "", err
	}

	party.Id = uuid.NewString()
	party.CreateTime = time.Now().Unix()

	m, err := mapper.ProtoToMap(party)
	if err != nil {
		return "", err
	}

	params := map[string]interface{}{
		"party": m,
	}

	_, err = session.Run(*cypher, params)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return party.Id, nil
}
