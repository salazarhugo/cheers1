package post

import (
	"encoding/json"
	"errors"
	"github.com/fatih/structs"
	"github.com/labstack/echo/v4"
	"github.com/lithammer/shortuuid/v4"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"log"
	"net/http"
	"rest-api/usereventpb"
	"rest-api/utils"
	"strconv"
	"time"
)

//func GetHistoryPosts(c echo.Context) error {
//cc := c.(*utils.CustomContext)
//session := utils.GetSession(cc.Driver)
//defer session.Close()
//
//pageSize, err := strconv.Atoi(cc.QueryParam("pageSize"))
//page, err := strconv.Atoi(cc.QueryParam("page"))
//skip := page * pageSize
//
//cypher := ` MATCH (u:User {id: $userId})-[:FOLLOWS*0..1]->(:User)-[:POSTED]->(p:Post)
//			WHERE p.created > datetime().epochMillis-24*60*60*1000 AND p.type = 'IMAGE'
//			WITH p SKIP $skip LIMIT $pageSize
//			RETURN [p.locationLongitude, p.locationLatitude], p.photos[0] as photoUrl`
//
//params := map[string]interface{}{
//	"userId": cc.Get("userId"),
//	"skip":   skip, "pageSize": pageSize,
//}
//
//result, err := session.Run(cypher, params)
//
//if err != nil {
//	log.Fatalf("Neo4j Error: %s", err)
//	return err
//}
//
//features := make([]utils.Feature, 0)
//
//for result.Next() {
//	res := result.Record().Values[0].([]interface{}) // [45.643, 34.32]
//	photoUrl := result.Record().Values[1].(string)
//	var coordinates []float64
//	coordinates = append(coordinates, res[0].(float64))
//	coordinates = append(coordinates, res[1].(float64))
//	features = append(features, utils.Feature{
//		Type:       "Feature",
//		Geometry:   utils.Geometry{Type: "Point", Coordinates: coordinates},
//		Properties: utils.Properties{PhotoUrl: photoUrl},
//	})
//}
//
//return cc.JSON(http.StatusOK, utils.FeatureCollection{Type: "FeatureCollection", Features: features.([]utils.Feature)})
//}

func GetMapPosts(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	pageSize, err := strconv.Atoi(cc.QueryParam("pageSize"))
	page, err := strconv.Atoi(cc.QueryParam("page"))
	skip := page * pageSize

	features, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			`MATCH (u:User {id: $userId})-[:FOLLOWS*0..1]->(:User)-[:POSTED]->(p:Post)
                    WHERE p.created > datetime().epochMillis-24*60*60*1000 AND p.type = 'IMAGE'
                    WITH p SKIP $skip LIMIT $pageSize
                    RETURN [p.locationLongitude, p.locationLatitude], p.photos[0] as photoUrl`,
			map[string]interface{}{"userId": cc.Get("userId"), "skip": skip, "pageSize": pageSize})
		if err != nil {
			log.Fatalf("Neo4j Error: %s", err)
			return nil, err
		}

		features := make([]utils.Feature, 0)

		for result.Next() {
			res := result.Record().Values[0].([]interface{}) // [45.643, 34.32]
			photoUrl := result.Record().Values[1].(string)
			var coordinates []float64
			coordinates = append(coordinates, res[0].(float64))
			coordinates = append(coordinates, res[1].(float64))
			features = append(features, utils.Feature{
				Type:       "Feature",
				Geometry:   utils.Geometry{Type: "Point", Coordinates: coordinates},
				Properties: utils.Properties{PhotoUrl: photoUrl},
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

func GetPosts(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	params := map[string]interface{}{
		"userId": cc.Get("userId"),
	}

	post, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			`MATCH (u:User { id: $userId })-[:POSTED]->(p:Post)
			RETURN p { .*, username: u.username, verified: u.verified, profilePictureUrl: u.profilePictureUrl }
			ORDER BY p.created DESC`,
			params,
		)
		if err != nil {
			return nil, err
		}

		posts := make([]interface{}, 0)
		for result.Next() {
			posts = append(posts, result.Record().Values[0])
		}

		return posts, nil
	})
	if err != nil {
		return err
	}

	return cc.JSON(http.StatusOK, post)
}

func PostMembers(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	postId := cc.Param("postId")
	pageSize, err := strconv.Atoi(cc.QueryParam("pageSize"))
	page, err := strconv.Atoi(cc.QueryParam("page"))
	skip := page * pageSize

	cypher := `
			MATCH (u:User {id: $userId})
			MATCH (p:Post {id: $postId})-[:WITH]->(w:User)
			WITH 
				w,
				exists((u)-[:FOLLOWS]->(w)) as followBack 
			SKIP 
				$skip 
			LIMIT 
				$pageSize
			RETURN w {
				.id,
				.name,
				.username,
				.verified,
				.profilePictureUrl,
				followBack: followBack
			}`

	params := map[string]interface{}{
		"userId":   cc.Get("userId"),
		"postId":   postId,
		"skip":     skip,
		"pageSize": pageSize,
	}

	result, err := session.Run(cypher, params)

	if err != nil {
		return err
	}

	users := make([]interface{}, 0)

	for result.Next() {
		users = append(users, result.Record().Values[0])
	}

	return cc.JSON(http.StatusOK, users)
}

// GET /posts/:userId
func GetUserPosts(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	otherUserId := cc.Param("userId")
	pageSize, err := strconv.Atoi(cc.QueryParam("pageSize"))
	page, err := strconv.Atoi(cc.QueryParam("page"))
	skip := page * pageSize

	cypher := `
			MATCH (u:User {id: $userId})
			MATCH (a:User)-[:POSTED]->(p:Post)
			WHERE a.id = $otherUserId OR a.username = $otherUserId
			OPTIONAL MATCH (:User)-[l:LIKED]->(p) 
			WITH 
				p,
				a,
				exists((u)-[:LIKED]->(p)) as liked,
				count(DISTINCT l) as likes 
			SKIP 
				$skip 
			LIMIT 
				$pageSize
			RETURN p {
				.*,
				likes: likes,
				liked: liked,
				username: a.username,
				verified: a.verified,
				profilePictureUrl: a.profilePictureUrl
			}`

	params := map[string]interface{}{
		"userId":      cc.Get("userId"),
		"otherUserId": otherUserId,
		"skip":        skip,
		"pageSize":    pageSize,
	}

	result, err := session.Run(cypher, params)

	if err != nil {
		return err
	}

	posts := make([]interface{}, 0)

	for result.Next() {
		posts = append(posts, result.Record().Values[0])
	}

	return cc.JSON(http.StatusOK, posts)
}

func PostFeed(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	pageSize, err := strconv.Atoi(cc.QueryParam("pageSize"))
	page, err := strconv.Atoi(cc.QueryParam("page"))
	skip := page * pageSize

	cypher := `MATCH (u:User {id: $userId})-[:FOLLOWS*0..1]->(a:User)-[:POSTED]->(p:Post)
			WITH u, a, p SKIP $skip LIMIT $pageSize
			OPTIONAL MATCH (:User)-[r:LIKED]->(p) 
			WITH p, a, exists((u)-[:LIKED]->(p)) as liked, count(DISTINCT r) as likes 
			RETURN p {
				.*,
				likes: likes,
				liked: liked,
				username: a.username,
				verified: a.verified,
				profilePictureUrl: a.profilePictureUrl
			} 
			ORDER BY p.created DESC`

	params := map[string]interface{}{
		"userId":   cc.Get("userId"),
		"skip":     skip,
		"pageSize": pageSize,
	}

	posts, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			cypher,
			params,
		)
		if err != nil {
			return nil, err
		}

		posts := make([]interface{}, 0)
		for result.Next() {
			posts = append(posts, result.Record().Values[0])
		}

		return posts, nil
	})
	if err != nil {
		return err
	}

	return cc.JSON(http.StatusOK, posts)
}

type Post struct {
	Id                string   `json:"id" structs:"id"`
	Caption           string   `json:"caption" structs:"caption"`
	Beverage          string   `json:"beverage" structs:"beverage"`
	Privacy           string   `json:"privacy" structs:"privacy"`
	Photos            []string `json:"photos" structs:"photos"`
	Drunkenness       int      `json:"drunkenness" structs:"drunkenness"`
	VideoUrl          string   `json:"videoUrl" structs:"videoUrl"`
	VideoThumbnailUrl string   `json:"videoThumbnailUrl" structs:"videoThumbnailUrl"`
	Latitude          float64  `json:"latitude" structs:"latitude"`
	Longitude         float64  `json:"longitude" structs:"longitude"`
	Altitude          float64  `json:"altitude" structs:"altitude"`
	LocationName      string   `json:"locationName" structs:"locationName"`
	Type              string   `json:"type" structs:"type"`
	TagUsersId        []string `json:"tagUsersId" structs:"tagUsersId"`
}

// POST("/posts/create")
func CreatePost(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	userId := cc.Get("userId").(string)
	post := Post{}

	err := json.NewDecoder(c.Request().Body).Decode(&post)
	if err != nil {
		return err
	}

	post.Id = shortuuid.New()
	cypher := `MATCH (u:User { id: $userId}) 
                CREATE (p: Post $post)
                SET p += { authorId: $userId, created: datetime().epochMillis } 
                CREATE (u)-[:POSTED]->(p)
                WITH p UNWIND $tagUsersId as tagUserId 
                MATCH (u2:User {id: tagUserId}) 
                CREATE (p)-[:WITH]->(u2)`

	params := map[string]interface{}{
		"userId":     userId,
		"post":       structs.Map(post),
		"tagUsersId": post.TagUsersId,
	}

	_, err = session.Run(cypher, params)
	if err != nil {
		log.Println(err)
		return err
	}

	state := &usereventpb.UserEvent{
		Type:   usereventpb.UserEventType_CREATE_POST,
		UserId: userId,
		PostId: post.Id,
		Time:   time.Now().Unix(),
	}
	err = utils.PublishProtoMessages(state)
	if err != nil {
		return err
	}
	return cc.NoContent(http.StatusOK)
}

// POST("/posts/:postId/delete")
func DeletePost(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	postId := cc.Param("postId")

	cypher := `MATCH (p:Post { id: $postId }) 
                DETACH DELETE p`

	params := map[string]interface{}{
		"userId": cc.Get("userId"),
		"postId": postId,
	}

	_, err := session.Run(cypher, params)
	if err != nil {
		return err
	}

	return cc.NoContent(http.StatusOK)
}

// POST("/posts/:postId/like")
func LikePost(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	userId := cc.Get("userId").(string)
	postId := cc.Param("postId")

	cypher := `MATCH (author:User)-[:POSTED]->(p:Post { id: $postId})
				MATCH (u:User {id: $userId}) 
                MERGE (u)-[:LIKED]->(p)
                RETURN { photoUrl: u.profilePictureUrl, username: u.username, name: u.name, 
                tokens: author.registrationTokens, authorId: author.id, beverage: p.beverage }`

	params := map[string]interface{}{
		"userId": userId,
		"postId": postId,
	}

	result, err := session.Run(cypher, params)
	if err != nil {
		log.Println(err)
		return err
	}

	if !result.Next() {
		return errors.New("expected one record was 0 or greater than 1")
	}

	record := result.Record().Values[0].(map[string]interface{})
	authorId := record["authorId"].(string)
	state := &usereventpb.UserEvent{
		Type:        usereventpb.UserEventType_POST_LIKE,
		UserId:      userId,
		OtherUserId: authorId,
		PostId:      postId,
		Time:        time.Now().Unix(),
	}
	utils.PublishProtoMessages(state)

	return cc.NoContent(http.StatusOK)
}

// POST("/posts/:postId/unlike")
func UnlikePost(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	cypher := `MATCH (u:User { id: $userId })-[l:LIKED]->(p:Post { id: $postId }) 
				DELETE l`

	params := map[string]interface{}{
		"userId": cc.Get("userId"),
		"postId": cc.Param("postId"),
	}

	session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		transaction.Run(
			cypher,
			params,
		)
		return nil, nil
	})

	return cc.String(http.StatusOK, "OK!")
}
