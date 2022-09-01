package user

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"github.com/fatih/structs"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"log"
	"net/http"
	"rest-api/usereventpb"
	"rest-api/utils"
	"strconv"
	"strings"
	"time"
)

type Activity struct {
	Avatar     string `json:"avatar,omitempty"`
	PhotoUrl   string `json:"photoUrl,omitempty"`
	Username   string `json:"username,omitempty"`
	Verified   bool   `json:"verified,omitempty"`
	Time       int64  `json:"time,omitempty"`
	FollowBack bool   `json:"followBack,omitempty"`
	Type       string `json:"type,omitempty"`
}

// GET("users/activity")
func GetActivity(c echo.Context) error {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "cheers-a275e")
	if err != nil {
		panic(err)
	}
	userId := c.Get("userId").(string)

	docs, err := client.Doc("users/"+userId).Collection("activity").OrderBy("time", firestore.Desc).Limit(20).Documents(ctx).GetAll()

	activities := make([]Activity, 0)

	for _, doc := range docs {
		data := doc.Data()

		userId := data["userId"].(string)
		postId := data["postId"].(string)
		Type := data["type"].(string)
		time := data["time"].(int64)

		userCard := GetUserCard(c, userId)

		photoUrl := ""
		if postId != "" {
			res := getPost(c, postId)
			if res != nil {
				photoUrl = *res
			}
		}

		activity := Activity{
			Avatar:     userCard.Avatar,
			PhotoUrl:   photoUrl,
			Username:   userCard.Username,
			Verified:   userCard.Verified,
			Time:       time,
			FollowBack: userCard.FollowBack,
			Type:       Type,
		}

		activities = append(activities, activity)
	}

	return c.JSON(http.StatusOK, activities)
}

type UserCard struct {
	Avatar     string `structs:"avatar,omitempty"`
	Username   string `structs:"username,omitempty"`
	Verified   bool   `structs:"verified,omitempty"`
	FollowBack bool   `structs:"followBack,omitempty"`
}

func GetUserCard(c echo.Context, userId string) *UserCard {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	cypher := `MATCH (u:User) 
                WHERE u.id = $userIdOrUsername OR u.username = $userIdOrUsername 
                OPTIONAL MATCH (:User)-[:FOLLOWS]->(u)
                WITH u, exists((:User { id: $userId })-[:FOLLOWS]->(u)) as followBack
                RETURN u {
					.id,
					.name,
					.username,
					.verified,
					avatar: u.profilePictureUrl,
					followBack: followBack
				}`

	params := map[string]interface{}{
		"userId":           cc.Get("userId"),
		"userIdOrUsername": userId,
	}

	userCard := &UserCard{}

	result, err := session.Run(cypher, params)

	if err != nil {
		return nil
	}

	if result.Next() {
		inter := result.Record().Values[0]
		err = mapstructure.Decode(inter, &userCard)
	}

	return userCard
}

func getPost(c echo.Context, postId string) *string {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	cypher := `MATCH (p:Post {id: $postId}) 
				WHERE p.type = 'IMAGE'
                RETURN p.photos[0]`

	params := map[string]interface{}{
		"postId": postId,
	}

	result, err := session.Run(cypher, params)

	if err != nil {
		return nil
	}

	if result.Next() {
		firstPhoto := result.Record().Values[0].(string)
		return &firstPhoto
	}

	return nil
}

// POST("/unfollow")
func UnfollowUser(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	username := cc.QueryParam("username")

	cypher := `MATCH (u:User)-[r:FOLLOWS]->(u2:User)
                 WHERE u.id = $userId AND u2.username = $username
                 DELETE r`

	params := map[string]interface{}{
		"userId":   cc.Get("userId"),
		"username": username,
	}

	session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		transaction.Run(cypher, params)
		return nil, nil
	})

	return cc.String(http.StatusOK, "OK!")
}

type Notification struct {
	OtherUserId string   `structs:"otherUserId,omitempty"`
	Tokens      []string `structs:"tokens,omitempty"`
	Avatar      string   `structs:"avatar,omitempty"`
	Username    string   `structs:"username,omitempty"`
}

// POST("/users/:userId/block")
func BlockUser(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	otherUserId := cc.Param("userId")

	cypher := `MATCH (u:User)-[r:FOLLOWS]-(u2:User)
                WHERE u.id = $userId AND u2.id = $otherUserId
                DELETE r MERGE (u)-[:BLOCKED]->(u2)`

	params := map[string]interface{}{
		"userId":      cc.Get("userId"),
		"otherUserId": otherUserId,
	}

	session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		transaction.Run(cypher, params)
		return nil, nil
	})

	return cc.String(http.StatusOK, "OK!")
}

// POST("/follow")
func FollowUser(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	username := cc.QueryParam("username")

	cypher := `MATCH (u:User), (f:User) 
                 WHERE u.id = $userId AND f.username = $username 
                 MERGE (u)-[:FOLLOWS]->(f) 
                 RETURN { avatar: u.profilePictureUrl, username: u.username, tokens: f.registrationTokens, otherUserId: f.id }`

	params := map[string]interface{}{
		"userId":   cc.Get("userId"),
		"username": username,
	}

	session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(cypher, params)
		if err != nil {
			panic(err)
		}

		notif := &Notification{}

		record, err := result.Single()
		err = mapstructure.Decode(record.Values[0].(map[string]interface{}), &notif)
		if err != nil {
			return nil, err
		}

		state := &usereventpb.UserEvent{
			Type:        usereventpb.UserEventType_FOLLOW,
			UserId:      cc.Get("userId").(string),
			OtherUserId: notif.OtherUserId,
			Time:        time.Now().Unix(),
		}
		utils.PublishProtoMessages(state)
		return nil, nil
	})

	return cc.String(http.StatusOK, "OK!")
}

// GET("users/available/:username")
func IsUsernameAvailable(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	username := cc.Param("username")

	user, err := session.ReadTransaction(matchUserNodeTxFunc(username))
	if err != nil {
		log.Fatalf("Error: %s", err)
		return err
	}

	return cc.JSON(http.StatusOK, user == nil)
}

// GET("users/:userIdOrUsername")
func GetUser(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	userIdOrUsername := cc.Param("userIdOrUsername")

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
				} as users`

	params := map[string]interface{}{
		"userId":           cc.Get("userId"),
		"userIdOrUsername": userIdOrUsername,
	}

	user, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(cypher, params)
		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().Values[0], nil
		}

		return nil, nil
	})

	if err != nil {
		log.Fatalf("Error: %s", err)
		return err
	}

	return cc.JSON(http.StatusOK, user)
}

// GET("users/suggestions")
func UserSuggestions(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	cypher := ` MATCH (u:User { id: $userId})-[:FOLLOWS]->(f:User)-[:FOLLOWS]->(suggestion:User)
				WHERE suggestion.id <> $userId
				AND NOT (u)-[:FOLLOWS]->(suggestion)
				WITH suggestion SKIP 0 LIMIT 10
                RETURN suggestion {
					.id,
					.name,
					.username,
					.verified,
					avatar: suggestion.profilePictureUrl
				}`

	params := map[string]interface{}{
		"userId": cc.Get("userId"),
	}

	result, err := session.Run(cypher, params)

	users := make([]interface{}, 0)

	for result.Next() {
		users = append(users, result.Record().Values[0])
	}

	if err = result.Err(); err != nil {
		log.Fatalf("Error: %s", err)
		return err
	}

	return cc.JSON(http.StatusOK, users)
}

// GetUserTickets GET("users/tickets")
func GetUserTickets(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	userId := cc.Get("userId").(string)
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "cheers-a275e")
	if err != nil {
		return cc.JSON(http.StatusInternalServerError, err.Error())
	}

	docs, err := client.Collection("stripe_customers").Doc(userId).Collection("tickets").Documents(ctx).GetAll()

	if err != nil {
		return cc.JSON(http.StatusInternalServerError, err.Error())
	}

	tickets := make([]interface{}, 0, 0)

	for _, doc := range docs {
		tickets = append(tickets, doc.Data())
	}

	return cc.JSON(http.StatusOK, tickets)
}

// SearchUsers GET("users/search/:query")
func SearchUsers(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	query := strings.ToLower(cc.Param("query"))
	if query == "" {
		return nil
	}

	cypher := `MATCH (u:User)
                WHERE u.username CONTAINS $query OR toLower(u.name) CONTAINS $query 
                AND u.id <> $userId
                WITH u LIMIT 10
                OPTIONAL MATCH (u)-[posts:POSTED]->(:Post)
                OPTIONAL MATCH (u)-[following:FOLLOWS]->(:User)
                OPTIONAL MATCH (:User)-[followers:FOLLOWS]->(u)
                WITH u, count(DISTINCT posts) as posts, count(DISTINCT following) as following, count(DISTINCT followers) as followers, exists((:User {id: $userId})-[:FOLLOWS]->(u)) as followBack
                RETURN u {.*, postCount: posts, followBack: followBack, following: following, followers: followers} as users`

	params := map[string]interface{}{
		"userId": cc.Get("userId"),
		"query":  query,
	}

	users, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(cypher, params)
		if err != nil {
			log.Fatalf("Neo4j Error: %s", err)
			return nil, err
		}

		users := make([]interface{}, 0)

		for result.Next() {
			users = append(users, result.Record().Values[0])
		}

		if err = result.Err(); err != nil {
			panic(err)
		}

		return users, nil
	})
	if err != nil {
		log.Fatalf("Error: %s", err)
		return err
	}

	return cc.JSON(http.StatusOK, users)
}

type User struct {
	Id                string `json:"id" structs:"id"`
	Username          string `json:"username" structs:"username"`
	Email             string `json:"email" structs:"email"`
	Name              string `json:"name" structs:"name"`
	ProfilePictureUrl string `json:"profilePictureUrl" structs:"profilePictureUrl"`
	Bio               string `json:"bio" structs:"bio"`
	Website           string `json:"website" structs:"website"`
	PhoneNumber       string `json:"phoneNumber" structs:"phoneNumber"`
}

func setRegistrationToken(userId string, registrationTokens []string) error {
	session := utils.GetSession(utils.GetDriver())
	defer session.Close()

	cypher := `MATCH (u:User {id: $userId})
				SET u.registrationTokens = $registrationTokens`

	params := map[string]interface{}{
		"userId":             userId,
		"registrationTokens": registrationTokens,
	}

	_, err := session.Run(cypher, params)

	return err
}

// POST("/users/tokens/:token")
func AddRegistrationToken(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	token := cc.Param("token")

	cypher := `MATCH (u:User {id: $userId})
				WHERE NOT $token IN u.registrationTokens OR u.registrationTokens IS NULL
				SET u.registrationTokens = coalesce(u.registrationTokens, []) + $token`

	params := map[string]interface{}{
		"userId": cc.Get("userId"),
		"token":  token,
	}

	_, err := session.Run(cypher, params)

	return err
}

// UpdateUser PATCH /users
func UpdateUser(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	userId := cc.Get("userId").(string)

	type UpdateUser struct {
		Email             *string `json:"email" structs:"email,omitempty"`
		Name              *string `json:"name" structs:"name,omitempty"`
		ProfilePictureUrl *string `json:"profilePictureUrl" structs:"profilePictureUrl,omitempty"`
		Bio               *string `json:"bio" structs:"bio,omitempty"`
		Website           *string `json:"website" structs:"website,omitempty"`
		PhoneNumber       *string `json:"phoneNumber" structs:"phoneNumber,omitempty"`
	}
	user := &UpdateUser{}

	err := json.NewDecoder(cc.Request().Body).Decode(&user)
	if err != nil {
		log.Println("Error decoding request body.")
		return cc.NoContent(http.StatusBadRequest)
	}

	cypher := ` MATCH (u: User {id: $userId})
                SET u += $user`

	params := map[string]interface{}{
		"userId": userId,
		"user":   structs.Map(user),
	}

	_, err = session.Run(cypher, params)

	if err != nil {
		log.Println(err)
		return err
	}

	return cc.NoContent(http.StatusOK)
}

// POST("/users/create")
func CreateUser(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	userReq := User{}
	userId := cc.Get("userId").(string)
	err := json.NewDecoder(cc.Request().Body).Decode(&userReq)

	if err != nil {
		return err
	}

	userReq.Id = userId
	if _, err := session.WriteTransaction(addUserNodeTxFunc(userReq)); err != nil {
		log.Print(err)
		return err
	}

	var user interface{}
	if user, err = session.ReadTransaction(matchUserNodeTxFunc(userReq.Id)); err != nil {
		log.Print(err)
		return err
	}

	return cc.JSON(http.StatusOK, user)
}

func addUserNodeTxFunc(userReq User) neo4j.TransactionWork {
	cypher := ` CREATE (u: User $user)
                SET u += { created: datetime().epochMillis, verified: false }`

	params := map[string]interface{}{
		"user": structs.Map(userReq),
	}
	return func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run(
			cypher,
			params,
		)
		if err != nil {
			return nil, err
		}

		return result.Consume()
	}
}

func matchUserNodeTxFunc(userIdOrUsername string) neo4j.TransactionWork {
	return func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run(
			`MATCH (u:User)
					WHERE u.id = $userIdOrUsername OR u.username = $userIdOrUsername
					RETURN u { .* }`,
			map[string]interface{}{
				"userIdOrUsername": userIdOrUsername,
			},
		)
		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().Values[0], nil
		}

		return nil, nil
	}
}

func UpdateLocation(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	longitude, err := strconv.ParseFloat(c.QueryParam("longitude"), 64)
	latitude, err := strconv.ParseFloat(c.QueryParam("latitude"), 64)

	params := map[string]interface{}{
		"userId":    cc.Get("userId"),
		"longitude": longitude,
		"latitude":  latitude,
	}

	features, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			`MATCH (u:User {id: $userId}) 
					SET u.longitude = $longitude, u.latitude = $latitude`,
			params,
		)
		if err != nil {
			log.Fatalf("Neo4j Error: %s", err)
			return nil, err
		}

		return result, nil
	})
	if err != nil {
		log.Fatalf("Error: %s", err)
		return err
	}
	log.Printf("features: %s", features)

	return c.HTML(http.StatusOK, "<html>OK</html>")
}

func GetLocations(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	params := map[string]interface{}{
		"userId": cc.Get("userId"),
	}

	features, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			`MATCH (u:User {id: $userId})-[:FOLLOWS]->(f:User)
			WHERE f.longitude IS NOT NULL AND f.latitude IS NOT NULL 
			WITH f LIMIT 100
			RETURN [f.longitude, f.latitude], f.username`,
			params,
		)

		if err != nil {
			log.Fatalf("Neo4j Error: %s", err)
			return nil, err
		}

		features := make([]utils.Feature, 0)

		for result.Next() {
			res := result.Record().Values[0].([]interface{}) // [45.643, 34.32]
			username := result.Record().Values[1].(string)
			var coordinates []float64
			coordinates = append(coordinates, res[0].(float64))
			coordinates = append(coordinates, res[1].(float64))
			features = append(features, utils.Feature{
				Type:       "Feature",
				Geometry:   utils.Geometry{Type: "Point", Coordinates: coordinates},
				Properties: utils.Properties{Username: username},
			})
		}

		return features, nil
	})

	if err != nil {
		log.Fatalf("Error: %s", err)
		return err
	}
	log.Printf("features: %s", features)

	return cc.JSON(http.StatusOK, utils.FeatureCollection{Type: "FeatureCollection", Features: features.([]utils.Feature)})
}

// GET("following/list")
func FollowingList(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	userId := cc.Get("userId")
	userIdParam := cc.QueryParam("userIdOrUsername")

	cypher := `
				MATCH (me: User {id: $userId})
				MATCH (u:User)-[:FOLLOWS]->(following:User)
				WHERE u.id = $userIdOrUsername OR u.username = $userIdOrUsername 
                OPTIONAL MATCH (following)-[:POSTED]->(s:Story) WHERE s.created > datetime().epochMillis-24*60*60*1000
                OPTIONAL MATCH (me)-[seen:SEEN]->(s)
                WITH 
					following,
					exists((me)-[:FOLLOWS]->(following)) as followBack,
					count(s) > 0 as hasStory,
					count(seen) = count(s) as seenStory
				LIMIT 20
                RETURN following { 
					.id,
					.name,
					.username,
					.verified,
					.profilePictureUrl,
					followBack: followBack,
					storyState:
						CASE 
							WHEN hasStory AND NOT seenStory THEN "NOT_SEEN"
							WHEN hasStory AND seenStory THEN "SEEN"
							ELSE "EMPTY"
						END
				}`

	params := map[string]interface{}{
		"userId":           userId,
		"userIdOrUsername": userIdParam,
	}

	users, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(cypher, params)
		if err != nil {
			log.Fatalf("Neo4j Error: %s", err)
			return nil, err
		}

		users := make([]interface{}, 0)

		for result.Next() {
			users = append(users, result.Record().Values[0])
		}

		if err = result.Err(); err != nil {
			panic(err)
		}

		return users, nil
	})
	if err != nil {
		log.Fatalf("Error: %s", err)
		return err
	}

	return cc.JSON(http.StatusOK, users)
}

// GET("followers/list")
func FollowersList(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	userId := cc.Get("userId")
	userIdParam := cc.QueryParam("userIdOrUsername")

	cypher := ` 
				MATCH (me:User {id: $userId})
				MATCH (u:User)<-[:FOLLOWS]-(follower:User)
				WHERE u.id = $userIdOrUsername OR u.username = $userIdOrUsername 
                OPTIONAL MATCH (follower)-[:POSTED]->(s:Story) WHERE s.created > datetime().epochMillis-24*60*60*1000
                OPTIONAL MATCH (me)-[seen:SEEN]->(s)
                WITH 
					follower,
					exists((u)-[:FOLLOWS]->(follower)) as followBack,
					count(s) > 0 as hasStory,
					count(seen) = count(s) as seenStory
				LIMIT 20
                RETURN 
					follower {
						.id,
						.name,
						.username,
						.verified,
						.profilePictureUrl,
						followBack: followBack,
						storyState:
							CASE 
								WHEN hasStory AND NOT seenStory THEN "NOT_SEEN"
								WHEN hasStory AND seenStory THEN "SEEN"
								ELSE "EMPTY"
							END
					}`

	params := map[string]interface{}{
		"userId":           userId,
		"userIdOrUsername": userIdParam,
	}

	users, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(cypher, params)
		if err != nil {
			log.Fatalf("Neo4j Error: %s", err)
			return nil, err
		}

		users := make([]interface{}, 0)

		for result.Next() {
			users = append(users, result.Record().Values[0])
		}

		if err = result.Err(); err != nil {
			panic(err)
		}

		return users, nil
	})
	if err != nil {
		log.Fatalf("Error: %s", err)
		return err
	}

	return cc.JSON(http.StatusOK, users)
}
