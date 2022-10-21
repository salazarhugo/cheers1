package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	utils "github.com/salazarhugo/cheers1/libs/utils"
	"log"
)

func (p *storyRepository) DeleteStory(
	storyID string,
) error {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/DeleteStory.cql")
	if err != nil {
		return err
	}

	params := map[string]interface{}{
		"storyID": storyID,
	}

	_, err = session.Run(*cypher, params)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
