package app

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func UserTopicSub(w http.ResponseWriter, r *http.Request) {
	var m PubSubMessage
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("ioutil.ReadAll: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// byte slice unmarshalling handles base64 decoding.
	if err := json.Unmarshal(body, &m); err != nil {
		log.Printf("json.Unmarshal: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	//userevent := &user.UserEvent{}
	//err = proto.Unmarshal(m.Message.Data, userevent)
	//if err != nil {
	//	log.Fatal(err)
	//	return err
	//}
	//
	//switch userevent.Type {
	//case usereventpb.UserEventType_POST_LIKE:
	//	likePostNotification(c, userevent)
	//case usereventpb.UserEventType_CREATE_POST:
	//	postNotification(c, userevent)
	//case usereventpb.UserEventType_FOLLOW:
	//	followNotification(c, userevent)
	//case usereventpb.UserEventType_COMMENT:
	//	commentNotification(c, userevent)
	//default:
	//}
}
