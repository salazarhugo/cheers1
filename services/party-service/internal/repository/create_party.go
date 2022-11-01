package repository

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	party "github.com/salazarhugo/cheers1/genproto/cheers/type/party"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
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
	party.CreateTime = timestamppb.Now()

	bytes, err := protojson.Marshal(party)
	if err != nil {
		return nil, err
	}

	var m = make(map[string]interface{}, 0)

	err = json.Unmarshal(bytes, &m)
	if err != nil {
		log.Println(err)
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
