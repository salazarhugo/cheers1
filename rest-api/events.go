package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"log"
	"net/http"
	"strconv"
)

func updateEvent(c echo.Context) error {
	cc := c.(*CustomContext)
	session := getSession(cc.Driver)

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

func createEvent(c echo.Context) error {
	cc := c.(*CustomContext)
	session := getSession(cc.Driver)

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

	return cc.String(http.StatusOK, "OK!")
}

func uninterestEvent(c echo.Context) error {
	cc := c.(*CustomContext)
	session := getSession(cc.Driver)
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
func ungoingEvent(c echo.Context) error {
	cc := c.(*CustomContext)
	session := getSession(cc.Driver)
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
func goingEvent(c echo.Context) error {
	cc := c.(*CustomContext)
	session := getSession(cc.Driver)
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

func interestEvent(c echo.Context) error {
	cc := c.(*CustomContext)
	session := getSession(cc.Driver)
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

func getEvents(c echo.Context) error {
	cc := c.(*CustomContext)
	session := getSession(cc.Driver)
	defer session.Close()

	pageSize, err := strconv.Atoi(cc.QueryParam("pageSize"))
	page, err := strconv.Atoi(cc.QueryParam("page"))
	skip := page * pageSize

	events, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			`MATCH (u:User { userId: $userId })-[:POSTED]->(e:Event)
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

func getEventFeed(c echo.Context) error {
	cc := c.(*CustomContext)
	session := getSession(cc.Driver)
	defer session.Close()

	pageSize, err := strconv.Atoi(cc.QueryParam("pageSize"))
	page, err := strconv.Atoi(cc.QueryParam("page"))
	skip := page * pageSize

	events, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			`MATCH (host:User)-[:POSTED]->(e:Event)
                    WITH e SKIP $skip LIMIT $pageSize 
					MATCH (u: User {id: $userId})
                    OPTIONAL MATCH (:User)-[interest:INTERESTED]->(e)
                    OPTIONAL MATCH (:User)-[going:GOING]->(e)
                    WITH e, exists((u)-[:INTERESTED]->(e)) as interested, count(DISTINCT interest) as interestedCount,
					exists((u)-[:GOING]->(e)) as going, count(DISTINCT going) as goingCount
					RETURN e { .*, interested: interested, interestedCount: interestedCount, going: going, goingCount: goingCount } 
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

func inviteEvent(c echo.Context) error {
	cc := c.(*CustomContext)
	session := getSession(cc.Driver)
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

func deleteEvent(c echo.Context) error {
	cc := c.(*CustomContext)
	session := getSession(cc.Driver)
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
