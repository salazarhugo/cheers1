package repository

import (
	"cloud.google.com/go/datastore"
	"context"
	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	party "github.com/salazarhugo/cheers1/gen/go/cheers/type/party"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/mapper"
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

	m, err := mapper.ProtoToMap(party)
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

	err = PersistToDataStore(party)
	if err != nil {
		log.Println(err)
	}

	return party, nil
}

type Party struct {
	Name    string    `datastore:"name"`
	Created time.Time `datastore:"created"`
	Desc    string    `datastore:"description"`
	id      int64     // The integer ID used in the datastore.
}

func PersistToDataStore(party2 *party.Party) error {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, "cheers-a275")
	if err != nil {
		log.Fatalf("Could not create datastore client: %v", err)
		return err
	}
	defer client.Close()

	data := &Party{
		Desc:    party2.Description,
		Name:    party2.Name,
		Created: time.Now(),
	}
	key := datastore.IncompleteKey("Party", nil)
	_, err = client.Put(ctx, key, data)

	return err
}
