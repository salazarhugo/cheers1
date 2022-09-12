package event

import (
	"encoding/json"
	"github.com/fatih/structs"
	"github.com/labstack/echo/v4"
	"github.com/lithammer/shortuuid/v4"
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

type Party struct {
	Id           string  `json:"id" structs:"id"`
	Name         string  `json:"name" structs:"name"`
	Description  string  `json:"description" structs:"description"`
	Address      string  `json:"address" structs:"address"`
	Privacy      string  `json:"privacy" structs:"privacy"`
	BannerUrl    string  `json:"bannerUrl" structs:"bannerUrl"`
	StartDate    int64   `json:"startDate" structs:"startDate"`
	EndDate      int64   `json:"endDate" structs:"endDate"`
	HostId       string  `json:"hostId" structs:"hostId"`
	Longitude    float64 `json:"longitude" structs:"longitude"`
	Latitude     float64 `json:"latitude" structs:"latitude"`
	LocationName string  `json:"locationName" structs:"locationName"`
}

func CreateParty(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)

	party := Party{}

	err := json.NewDecoder(c.Request().Body).Decode(&party)
	if err != nil {
		log.Println(err)
		return err
	}

	party.Id = shortuuid.New()

	cypher := `MATCH (u:User { id: $userId}) 
				CREATE (e: Party $party)
		   		SET e += { created: datetime().epochMillis } 
				CREATE (u)-[:POSTED]->(e)
				CREATE (u)-[:GOING]->(e)`

	params := map[string]interface{}{
		"userId": cc.Get("userId"),
		"party":  structs.Map(party),
	}

	_, err = session.Run(cypher, params)
	if err != nil {
		log.Println(err)
		return err
	}

	cypher = `
			MATCH (u:User {id: $userId})
			MATCH (u)-[:POSTED]->(party: Party {id: $partyId})
			RETURN 
			party {
				.*
			}`

	params = map[string]interface{}{
		"userId":  cc.Get("userId"),
		"partyId": party.Id,
	}

	result, err := session.Run(cypher, params)
	if err != nil {
		log.Println(err)
		return err
	}

	var response interface{}

	if result.Next() {
		response = result.Record().Values[0]
	}

	return cc.JSON(http.StatusOK, response)
}

func UninterestParty(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	partyId := c.QueryParam("partyId")

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			`MATCH (u:User { id: $userId })-[l:INTERESTED]->(p:Party { id: $partyId }) 
                    DELETE l`,
			map[string]interface{}{"userId": c.Get("userId"), "partyId": partyId})
		if err != nil {
			return nil, err
		}

		return nil, result.Err()
	})

	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
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

func GetParty(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	partyId := cc.Param("partyId")

	cypher := `
			MATCH (u:User {id: $userId})
			MATCH (host:User)-[:POSTED]->(party: Party {id: $partyId})
			OPTIONAL MATCH (:User)-[interest:INTERESTED]->(party)
			OPTIONAL MATCH (:User)-[going:GOING]->(party)
			WITH 
				host,
				party, 
				exists((u)-[:INTERESTED]->(party)) as interested,
				exists((u)-[:GOING]->(party)) as going,
				count(DISTINCT interest) as interestedCount,
				count(DISTINCT going) as goingCount,
				host.id = u.id as isHost
			RETURN 
				party {
					.*,
					isHost: isHost,
					hostId: host.id,
					hostName: host.name,
					interested: interested,
					interestedCount: interestedCount,
					going: going,
					goingCount: goingCount
				}`

	params := map[string]interface{}{
		"userId":  cc.Get("userId"),
		"partyId": partyId,
	}

	result, err := session.Run(cypher, params)

	if err != nil {
		return err
	}

	if result.Next() {
		party := result.Record().Values[0]
		return cc.JSON(http.StatusOK, party)
	}

	return cc.NoContent(http.StatusOK)
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

func InterestParty(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	partyId := c.QueryParam("partyId")

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			`MATCH (p:Party { id: $partyId}), (u:User {id: $userId}) 
                MERGE (u)-[:INTERESTED]->(p)`,
			map[string]interface{}{"userId": c.Get("userId"), "partyId": partyId})
		if err != nil {
			return nil, err
		}

		return nil, result.Err()
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "")
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

func GetMyParties(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	pageSize, err := strconv.Atoi(cc.QueryParam("pageSize"))
	page, err := strconv.Atoi(cc.QueryParam("page"))
	skip := page * pageSize

	parties, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			`
				MATCH (u: User{id:$userId})
				OPTIONAL MATCH (u)-[:POSTED]->(p:Party)
				OPTIONAL MATCH (:User)-[interest:INTERESTED]->(p)
				OPTIONAL MATCH (:User)-[going:GOING]->(p)
				WITH 
					u, 
					p,
					exists((u)-[:INTERESTED]->(p)) as interested,
					count(DISTINCT interest) as interestedCount,
					exists((u)-[:GOING]->(p)) as going,
					count(DISTINCT going) as goingCount
				SKIP 
					$skip 
				LIMIT 
					$pageSize
				CALL {
					WITH u, p
					OPTIONAL MATCH (u)-[:FOLLOWS]->(mutual:User)-[:GOING]->(p)
					WITH mutual, count(mutual) as mutualCount LIMIT 2
					RETURN collect(mutual.picture) as mutualpictures,
					collect(mutual.username) as mutualUsernames, mutualCount
				}
				RETURN p { 
					.*, 
					hostId: u.id,
					hostName: u.name,
					interested: interested,
					interestedCount: interestedCount,
					going: going,
					goingCount: goingCount,
					mutualCount: mutualCount,
					mutualUsernames: mutualUsernames,
					mutualpictures: mutualpictures
				} 
				ORDER BY p.startDate DESC`,
			map[string]interface{}{"userId": cc.Get("userId"), "skip": skip, "pageSize": pageSize})
		if err != nil {
			log.Fatalf("Neo4j Error: %s", err)
			return nil, err
		}

		parties := make([]interface{}, 0)

		for result.Next() {
			parties = append(parties, result.Record().Values[0])
		}

		return parties, nil
	})
	if err != nil {
		panic(err)
	}

	return cc.JSON(http.StatusOK, parties)
}

func GetPartyFeed(c echo.Context) error {
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
				OPTIONAL MATCH (e1:Party {privacy: "PUBLIC"})
				OPTIONAL MATCH (u)-[:FOLLOWS]->(author:User)-[:POSTED]->(e2:Party {privacy: "FRIENDS"})
				OPTIONAL MATCH (e3:Party)-[:INVITED]->(u)
				UNWIND [val in [e1, e2, e3] WHERE val is not null] as e
				OPTIONAL MATCH (:User)-[interest:INTERESTED]->(e)
				OPTIONAL MATCH (:User)-[going:GOING]->(e)
				WITH 
					u, 
					e,
					author,
					exists((u)-[:INTERESTED]->(e)) as interested,
					count(DISTINCT interest) as interestedCount,
					exists((u)-[:GOING]->(e)) as going,
					count(DISTINCT going) as goingCount
				SKIP 
					$skip 
				LIMIT 
					$pageSize
				CALL {
					WITH u, e, author
					OPTIONAL MATCH (u)-[:FOLLOWS]->(mutual:User)-[:GOING]->(e)
					WITH mutual, count(mutual) as mutualCount LIMIT 2
					RETURN collect(mutual.picture) as mutualpictures,
					collect(mutual.username) as mutualUsernames, mutualCount
				}
				RETURN e { 
					.*, 
					hostId: author.id,
					hostName: author.name,
					interested: interested,
					interestedCount: interestedCount,
					going: going,
					goingCount: goingCount,
					mutualCount: mutualCount,
					mutualUsernames: mutualUsernames,
					mutualpictures: mutualpictures
				} 
				ORDER BY e.startDate DESC`,
			map[string]interface{}{"userId": cc.Get("userId"), "skip": skip, "pageSize": pageSize})
		if err != nil {
			log.Fatalf("Neo4j Error: %s", err)
			return nil, err
		}

		parties := make([]interface{}, 0)

		for result.Next() {
			parties = append(parties, result.Record().Values[0])
		}

		return parties, nil
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

func DeleteParty(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	partyId := cc.QueryParam("partyId")

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			`MATCH (e: Party { id: $partyId}) 
					DETACH DELETE e`,
			map[string]interface{}{"userId": cc.Get("userId"), "partyId": partyId})
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
			 	RETURN g { .id, .username, .verified, .picture, .name, followBack: followBack}`

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
			 	RETURN g { .id, .username, .verified, .picture, .name, followBack: followBack}`

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
