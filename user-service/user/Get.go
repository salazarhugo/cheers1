package user

import (
	"encoding/json"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"salazar/cheers/user/proto/userpb"
	"salazar/cheers/user/utils"
)

// GetUser GET /?username={username}
func GetUser(
	currentUserId string,
	userIdOrUsername string,
) error {
	driver := utils.GetDriver()
	session := utils.GetSession(driver)
	defer session.Close()

	if userIdOrUsername == "" {
		return nil
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
					count(DISTINCT authorPosts) as post_count,
					count(DISTINCT following) as following_count, 
                	count(DISTINCT followers) as followers_count,
					exists((me)-[:FOLLOWS]->(u)) as follow_back,
					count(s) > 0 as has_story,
					count(seen) = count(s) as seen_story
                RETURN 
					{
						user: PROPERTIES(u),
						post_count: post_count,
						follow_back: follow_back,
						following_count: following_count,
						followers_count: followers_count,
						story_state:
							CASE 
								WHEN has_story AND NOT seen_story THEN "NOT_SEEN"
								WHEN has_story AND seen_story THEN "SEEN"
								ELSE "EMPTY"
							END
					}`

	params := map[string]interface{}{
		"userId":           currentUserId,
		"userIdOrUsername": userIdOrUsername,
	}

	_, err := session.Run(cypher, params)

	//var user *userpb.GetUserResponse
	//if result.Next() {
	//	userMap := result.Record().Values[0].(map[string]interface{})
	//	user, err = MapToUser(userMap)
	//} else {
	//return cc.JSON(http.StatusNoContent, map[string]interface{}{
	//	"username": userIdOrUsername,
	//	"userId":   cc.Get("userId"),
	//})
	//}

	if err != nil {
		log.Fatalf("Error: %s", err)
		return err
	}

	//return cc.JSON(http.StatusOK, user)
	return nil
}

func MapToUser(m map[string]interface{}) (*userpb.GetUserResponse, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	s := &userpb.GetUserResponse{}
	err = protojson.Unmarshal(b, s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
