package story

import (
	"encoding/json"
	"github.com/fatih/structs"
	"github.com/labstack/echo/v4"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"log"
	"net/http"
	"rest-api/utils"
	"strconv"
)

func GetMapStories(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	pageSize, err := strconv.ParseInt(cc.QueryParam("pageSize"), 10, 64)

	features, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			`MATCH (u:User {id: $userId})-[:FOLLOWS*0..1]->(:User)-[:POSTED]->(s:Story)
                    WHERE s.created > datetime().epochMillis-24*60*60*1000
					AND s.longitude IS NOT NULL AND s.latitude IS NOT NULL 
                    WITH s LIMIT $pageSize
                    RETURN [s.longitude, s.latitude], s.photoUrl as photoUrl`,
			map[string]interface{}{"userId": cc.Get("userId"), "pageSize": pageSize})
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

// POST("/stories/:storyId/seen")
func SeenStory(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	storyId := cc.Param("storyId")

	cypher := `MATCH (s:Story {id: $storyId}), (u:User { id: $userId }) 
                MERGE (u)-[:SEEN]->(s) 
                SET s.seenBy = s.seenBy + $userId`

	params := map[string]interface{}{
		"userId":  cc.Get("userId"),
		"storyId": storyId,
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

func StoryFeed(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	pageSize, err := strconv.Atoi(cc.QueryParam("pageSize"))
	page, err := strconv.Atoi(cc.QueryParam("page"))
	skip := page * pageSize

	cypher := `MATCH (u:User {id: $userId})-[:FOLLOWS*0..1]->(author:User)-[:POSTED]->(s:Story)
				WHERE s.created > datetime().epochMillis-24*60*60*1000
				WITH author, s, exists((u)-[:SEEN]->(s)) as seen SKIP $skip LIMIT $pageSize
				RETURN s {
					.*,
					seen: seen,
					username: author.username,
					verified: author.verified,
					profilePictureUrl: author.profilePictureUrl
				}
				ORDER BY s.created DESC`

	params := map[string]interface{}{
		"userId":   cc.Get("userId"),
		"skip":     skip,
		"pageSize": pageSize,
	}

	stories, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			cypher,
			params,
		)
		if err != nil {
			return nil, err
		}

		stories := make([]interface{}, 0)
		for result.Next() {
			stories = append(stories, result.Record().Values[0])
		}

		return stories, nil
	})
	if err != nil {
		return err
	}

	return cc.JSON(http.StatusOK, stories)
}

// POST("/stories/:storyId/delete")
func DeleteStory(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	storyId := cc.Param("storyId")

	cypher := `MATCH (p:Story { id: $storyId }) 
                DETACH DELETE p`

	params := map[string]interface{}{
		"userId":  cc.Get("userId"),
		"storyId": storyId,
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

type Story struct {
	Privacy      string  `json:"privacy" structs:"privacy"`
	PhotoUrl     string  `json:"photoUrl" structs:"photoUrl" `
	VideoUrl     string  `json:"videoUrl" structs:"videoUrl"`
	Latitude     float64 `json:"latitude" structs:"latitude"`
	Longitude    float64 `json:"longitude" structs:"longitude"`
	Altitude     float64 `json:"altitude" structs:"altitude"`
	LocationName string  `json:"locationName" structs:"locationName"`
	Type         string  `json:"type" structs:"type"`
}

func CreateStory(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	story := Story{}

	err := json.NewDecoder(c.Request().Body).Decode(&story)
	if err != nil {
		return err
	}

	log.Println(story)

	cypher := `MATCH (u:User { id: $userId})
                CREATE (s: Story $story)
                SET s += { id: apoc.create.uuid(), created: datetime().epochMillis, authorId: $userId }
                CREATE (u)-[:POSTED]->(s)
                RETURN s.id`

	params := map[string]interface{}{
		"userId": cc.Get("userId"),
		"story":  structs.Map(story),
	}

	session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		_, err := transaction.Run(
			cypher,
			params,
		)

		if err != nil {
			log.Fatalf("Neo4j Error: %s", err)
			return nil, err
		}

		return nil, nil
	})

	return cc.String(http.StatusOK, "OK!")
}
