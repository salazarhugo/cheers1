package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"salazar/cheers/email/usereventpb"
)

func main() {
	http.HandleFunc("/", UserEventPubSub)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

type PubSubMessage struct {
	Message struct {
		Data []byte `json:"data,omitempty"`
		ID   string `json:"id"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}

func UserEventPubSub(w http.ResponseWriter, r *http.Request) {
	var m PubSubMessage
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("ioutil.ReadAll: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &m); err != nil {
		log.Printf("json.Unmarshal: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	userevent := &usereventpb.UserEvent{}
	err = proto.Unmarshal(m.Message.Data, userevent)
	if err != nil {
		log.Fatal(err)
		return
	}

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "cheers-a275e")
	if err != nil {
		panic(err)
	}

	event := map[string]interface{}{
		"type":        userevent.GetType().String(),
		"userId":      userevent.UserId,
		"otherUserId": userevent.OtherUserId,
		"postId":      userevent.PostId,
		"eventId":     userevent.EventId,
		"time":        userevent.Time,
	}

	_, _, err = client.Doc("users/"+userevent.OtherUserId).Collection("activity").Add(ctx, event)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("%s %s %s", userevent.UserId, userevent.Type.String(), userevent.EventId)
}
