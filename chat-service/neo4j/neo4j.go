package neo4j

import (
	"chat/chatpb"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func GetUsers(userIds []string) ([]*chatpb.UserCard, error) {
	driver, err := neo4j.NewDriver(
		"neo4j+s://528a253a.databases.neo4j.io:7687",
		neo4j.BasicAuth("neo4j", "XRoQ6Lmz9QlFFTcwCWIWwR1o88hLfzV_HnP9mzDJuwc", ""))

	if err != nil {
		return nil, err
	}

	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher := ` MATCH (u:User)
				WHERE u.id IN $userIds
                RETURN u {
					.id,
					.name,
					.username,
					.verified,
					avatar: u.profilePictureUrl
				}`

	params := map[string]interface{}{
		"userIds": userIds,
	}

	result, err := session.Run(cypher, params)

	users := make([]*chatpb.UserCard, 0)

	for result.Next() {
		userMap := result.Record().Values[0].(map[string]interface{})
		user := &chatpb.UserCard{
			Id:       userMap["id"].(string),
			Name:     userMap["name"].(string),
			Username: userMap["username"].(string),
			Verified: userMap["verified"].(bool),
			Avatar:   userMap["avatar"].(string),
		}
		users = append(users, user)
	}

	if err = result.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
