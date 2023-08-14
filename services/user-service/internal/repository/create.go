package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/mapper"
	"log"
	"time"
)

func (p *userRepository) CreateUser(
	userID string,
	username string,
	name string,
	picture string,
	email string,
) (string, error) {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/CreateUser.cql")
	if err != nil {
		return "", err
	}

	user := &user.User{}
	user.Id = userID
	user.CreateTime = time.Now().Unix()
	user.Username = username
	user.Name = name
	user.Picture = picture
	user.Email = email

	m, err := mapper.ProtoToMap(user)
	if err != nil {
		log.Println(err)
		return "", err
	}

	params := map[string]interface{}{
		"user": m,
	}

	log.Println(user)
	result, err := session.Run(*cypher, params)
	if err != nil {
		log.Println(err)
		return "", err
	}

	if result.Err() != nil {
		return "", result.Err()
	}

	return user.Id, nil
}
