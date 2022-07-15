package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"salazar/cheers/notification/usereventpb"
)

type CustomContext struct {
	echo.Context
	neo4j.Driver
}

func getDriver() neo4j.Driver {
	driver, err := neo4j.NewDriver(
		"neo4j+s://528a253a.databases.neo4j.io:7687",
		neo4j.BasicAuth("neo4j", "XRoQ6Lmz9QlFFTcwCWIWwR1o88hLfzV_HnP9mzDJuwc", ""))
	if err != nil {
		panic(err)
	}

	return driver
}

func getSession(driver neo4j.Driver) neo4j.Session {
	return driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
}

func main() {
	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c, getDriver()}
			return next(cc)
		}
	})

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://web-cheers.web.app", "http://localhost:4200"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	e.POST("/", HelloPubSub)
	e.Logger.Fatal(e.Start(":" + port))
}

type PubSubMessage struct {
	Message struct {
		Data []byte `json:"data,omitempty"`
		ID   string `json:"id"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}

// HelloPubSub receives and processes a Pub/Sub push message.
func HelloPubSub(c echo.Context) error {
	var m PubSubMessage
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("ioutil.ReadAll: %v", err)
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if err := json.Unmarshal(body, &m); err != nil {
		log.Printf("json.Unmarshal: %v", err)
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	userevent := &usereventpb.UserEvent{}
	err = proto.Unmarshal(m.Message.Data, userevent)
	if err != nil {
		log.Fatal(err)
		return err
	}

	switch userevent.Type {
	case usereventpb.UserEventType_POST_LIKE:
		likePostNotification(c, userevent)
	case usereventpb.UserEventType_CREATE_POST:
		postNotification(c, userevent)
	case usereventpb.UserEventType_FOLLOW:
		followNotification(c, userevent)
	case usereventpb.UserEventType_COMMENT:
		commentNotification(c, userevent)
	default:
	}

	return c.NoContent(http.StatusOK)
}
