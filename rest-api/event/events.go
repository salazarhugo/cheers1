package event

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"log"
	"net/http"
	"rest-api/utils"
	"strconv"
)

func UpdateEvent(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)

	event := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&event)
	if err != nil {
		return err
	}

	_, err = session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			`MATCH (e:Event { id: $eventId}) 
		   			SET e = $event`,
			map[string]interface{}{"userId": cc.Get("userId"), "eventId": event["id"], "event": event})
		if err != nil {
			log.Fatalf("Neo4j Error: %s", err)
			return nil, err
		}

		return nil, result.Err()
	})

	if err != nil {
		log.Fatalf("Error: %s", err)
		return err
	}

	return cc.String(http.StatusOK, "OK!")
}

func CreateEvent(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)

	event := make(map[string]interface{})

	err := json.NewDecoder(c.Request().Body).Decode(&event)
	if err != nil {
		return err
	}

	_, err = session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			`MATCH (u:User { id: $userId}) 
				CREATE (e: Event $event)
		   		SET e += { created: datetime().epochMillis } 
				CREATE (u)-[:POSTED]->(e)`,
			map[string]interface{}{"userId": cc.Get("userId"), "event": event})
		if err != nil {
			log.Fatalf("Neo4j Error: %s", err)
			return nil, err
		}

		return nil, result.Err()
	})

	if err != nil {
		log.Fatalf("Error: %s", err)
		return err
	}

	return cc.NoContent(http.StatusOK)
}

func UninterestEvent(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	eventId := c.QueryParam("eventId")

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			`MATCH (u:User { id: $userId })-[l:INTERESTED]->(e:Event { id: $eventId }) 
                    DELETE l`,
			map[string]interface{}{"userId": c.Get("userId"), "eventId": eventId})
		if err != nil {
			return nil, err
		}

		return nil, result.Err()
	})

	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "OK!")
}

// POST("/event/:eventId/ungoing")
func UngoingEvent(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	eventId := c.Param("eventId")

	cypher := `MATCH (u:User { id: $userId })-[l:GOING]->(e:Event { id: $eventId }) 
				DELETE l`

	params := map[string]interface{}{
		"userId":  c.Get("userId"),
		"eventId": eventId,
	}

	_, err := session.Run(cypher, params)

	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

// POST("/event/:eventId/going")
func GoingEvent(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	eventId := c.Param("eventId")

	cypher := `MATCH (e:Event { id: $eventId}), (u:User {id: $userId}) 
                MERGE (u)-[:GOING]->(e)`

	params := map[string]interface{}{
		"userId":  c.Get("userId"),
		"eventId": eventId,
	}

	_, err := session.Run(cypher, params)

	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func InterestEvent(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	eventId := c.QueryParam("eventId")

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			`MATCH (e:Event { id: $eventId}), (u:User {id: $userId}) 
                MERGE (u)-[:INTERESTED]->(e)`,
			map[string]interface{}{"userId": c.Get("userId"), "eventId": eventId})
		if err != nil {
			return nil, err
		}

		return nil, result.Err()
	})

	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "OK!")
}

func GetEvents(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	pageSize, err := strconv.Atoi(cc.QueryParam("pageSize"))
	page, err := strconv.Atoi(cc.QueryParam("page"))
	skip := page * pageSize

	events, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			`MATCH (u:User { id: $userId })-[:POSTED]->(e:Event)
                    WITH u, e SKIP $skip LIMIT $pageSize 
                    OPTIONAL MATCH (:User)-[interest:INTERESTED]->(e)
                    OPTIONAL MATCH (:User)-[going:GOING]->(e)
                    WITH e, exists((u)-[:INTERESTED]->(e)) as interested, count(DISTINCT interest) as interestedCount, count(DISTINCT going) as goingCount
					RETURN e { .*, interested: interested, interestedCount: interestedCount, goingCount: goingCount } 
					ORDER BY e.startDate DESC`,
			map[string]interface{}{"userId": cc.Get("userId"), "skip": skip, "pageSize": pageSize})
		if err != nil {
			log.Fatalf("Neo4j Error: %s", err)
			return nil, err
		}

		if result.Err() != nil {
			log.Printf("GetEvents: result.Err() %v", result.Err())
			return nil, result.Err()
		}

		events := make([]interface{}, 0)

		for result.Next() {
			events = append(events, result.Record().Values[0])
		}

		return events, nil
	})
	if err != nil {
		panic(err)
	}

	return cc.JSON(http.StatusOK, events)
}

func GetEventFeed(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	pageSize, err := strconv.Atoi(cc.QueryParam("pageSize"))
	page, err := strconv.Atoi(cc.QueryParam("page"))
	skip := page * pageSize

	events, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			`
				MATCH (u: User{id:$userId})
				OPTIONAL MATCH (e1:Event {privacy: "PUBLIC"})
				OPTIONAL MATCH (u)-[:FOLLOWS]->(:User)-[:POSTED]->(e2:Event {privacy: "FRIENDS"})
				OPTIONAL MATCH (e3:Event)-[:INVITED]->(u)
				UNWIND [val in [e1, e2, e3] WHERE val is not null] as e
				OPTIONAL MATCH (:User)-[interest:INTERESTED]->(e)
				OPTIONAL MATCH (:User)-[going:GOING]->(e)
				WITH 
					u, 
					e,
					exists((u)-[:INTERESTED]->(e)) as interested,
					count(DISTINCT interest) as interestedCount,
					exists((u)-[:GOING]->(e)) as going,
					count(DISTINCT going) as goingCount
				SKIP 
					$skip 
				LIMIT 
					$pageSize
				CALL {
					WITH u, e
					OPTIONAL MATCH (u)-[:FOLLOWS]->(mutual:User)-[:GOING]->(e)
					WITH mutual, count(mutual) as mutualCount LIMIT 2
					RETURN collect(mutual.profilePictureUrl) as mutualProfilePictureUrls,
					collect(mutual.username) as mutualUsernames, mutualCount
				}
				RETURN e { 
					.*, 
					interested: interested,
					interestedCount: interestedCount,
					going: going,
					goingCount: goingCount,
					mutualCount: mutualCount,
					mutualUsernames: mutualUsernames,
					mutualProfilePictureUrls: mutualProfilePictureUrls
				} 
				ORDER BY e.startDate DESC`,
			map[string]interface{}{"userId": cc.Get("userId"), "skip": skip, "pageSize": pageSize})
		if err != nil {
			log.Fatalf("Neo4j Error: %s", err)
			return nil, err
		}

		events := make([]interface{}, 0)

		for result.Next() {
			events = append(events, result.Record().Values[0])
		}

		return events, nil
	})
	if err != nil {
		panic(err)
	}

	return cc.JSON(http.StatusOK, events)
}

func InviteEvent(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	eventId := cc.QueryParam("eventId")
	guests := cc.QueryParam("users")

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			`MATCH (e: Event { id: $eventId}) 
					MATCH (u: User {id:$userId})
                	WITH p UNWIND $guests as guest 
                	MATCH (u2:User {id: guest}) 
					CREATE (u)-[:INVITE]->(u2)`,
			map[string]interface{}{
				"userId":  cc.Get("userId"),
				"eventId": eventId,
				"guests":  guests,
			},
		)
		if err != nil {
			log.Fatalf("Neo4j Error: %s", err)
			return nil, err
		}

		return nil, result.Err()
	})

	if err != nil {
		log.Fatalf("Error: %s", err)
		return err
	}

	return cc.String(http.StatusOK, "OK!")
}

func DeleteEvent(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	eventId := cc.QueryParam("eventId")

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			`MATCH (e: Event { id: $eventId}) 
					DETACH DELETE e`,
			map[string]interface{}{"userId": cc.Get("userId"), "eventId": eventId})
		if err != nil {
			log.Fatalf("Neo4j Error: %s", err)
			return nil, err
		}

		return nil, result.Err()
	})

	if err != nil {
		log.Fatalf("Error: %s", err)
		return err
	}

	return cc.String(http.StatusOK, "OK!")
}

// GET /event/:eventId/interested/list

func InterestedList(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	userId := cc.Get("userId")

	eventId := cc.Param("eventId")

	cypher := `MATCH (e:Event {id: $eventId})<-[:INTERESTED]-(g:User)
				MATCH (u: User {id:$userId}) 
			 	WITH g, exists((u)-[:FOLLOWS]->(g)) as followBack LIMIT 20
			 	RETURN g { .id, .username, .verified, .profilePictureUrl, .name, followBack: followBack}`

	params := map[string]interface{}{
		"userId":  userId,
		"eventId": eventId,
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

// GET /event/:eventId/going/list

func GoingList(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	userId := cc.Get("userId")

	eventId := cc.Param("eventId")

	cypher := `MATCH (e:Event {id: $eventId})<-[:GOING]-(g:User)
				MATCH (u: User {id:$userId}) 
			 	WITH g, exists((u)-[:FOLLOWS]->(g)) as followBack LIMIT 20
			 	RETURN g { .id, .username, .verified, .profilePictureUrl, .name, followBack: followBack}`

	params := map[string]interface{}{
		"userId":  userId,
		"eventId": eventId,
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
