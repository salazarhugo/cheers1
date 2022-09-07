package user

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net/http"
	"salazar/cheers/user/userpb"
	"salazar/cheers/user/utils"
)

// GetUser GET /?username={username}
func GetUser(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	userIdOrUsername := cc.QueryParam("username")
	if userIdOrUsername == "" {
		return cc.JSON(http.StatusBadRequest, "Missing parameter: username")
	}

	cypher := `
				MATCH (me:User {id: $userId}) 
				MATCH (u:User) 
                WHERE u.id = $userIdOrUsername OR u.username = $userIdOrUsername 
                OPTIONAL MATCH (u)-[authorPosts:POSTED]->(:Post) 
                OPTIONAL MATCH (u)-[following:FOLLOWS]->(:User) 
                OPTIONAL MATCH (:User)-[followers:FOLLOWS]->(u)
                OPTIONAL MATCH (u)-[:POSTED]->(s:Story) WHERE s.created > datetime().epochMillis-24*60*60*1000
                OPTIONAL MATCH (me)-[seen:SEEN]->(s)
                WITH 
					u,
					count(DISTINCT authorPosts) as postCount,
					count(DISTINCT following) as following, 
                	count(DISTINCT followers) as followers,
					exists((me)-[:FOLLOWS]->(u)) as followBack,
					count(s) > 0 as hasStory,
					count(seen) = count(s) as seenStory
                RETURN u { 
					.*,
					postCount: postCount,
					followBack: followBack,
					following: following,
					followers: followers,
					storyState:
						CASE 
							WHEN hasStory AND NOT seenStory THEN "NOT_SEEN"
							WHEN hasStory AND seenStory THEN "SEEN"
							ELSE "EMPTY"
						END
				}`

	params := map[string]interface{}{
		"userId":           cc.Get("userId"),
		"userIdOrUsername": userIdOrUsername,
	}

	result, err := session.Run(cypher, params)

	var user *userpb.User
	if result.Next() {
		userMap := result.Record().Values[0].(map[string]interface{})
		user, err = MapToUser(userMap)
	} else {
		return cc.JSON(http.StatusNoContent, map[string]interface{}{
			"username": userIdOrUsername,
			"userId":   cc.Get("userId"),
		})
	}

	if err != nil {
		log.Fatalf("Error: %s", err)
		return err
	}

	return cc.JSON(http.StatusOK, user)
}

func MapToUser(m map[string]interface{}) (*userpb.User, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	s := &userpb.User{}
	err = protojson.Unmarshal(b, s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
