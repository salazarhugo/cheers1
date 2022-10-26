package app

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"net/http"
)

// PubSubMessage is the payload of a Pub/Sub event.
// See the documentation for more details:
// https://cloud.google.com/pubsub/docs/reference/rest/v1/PubsubMessage
type PubSubMessage struct {
	Message struct {
		Data []byte `json:"data,omitempty"`
		ID   string `json:"id"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}

func UserEventPubSub(c echo.Context) error {
	var m PubSubMessage
	body, err := io.ReadAll(c.Request().Body)
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
