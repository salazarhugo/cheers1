package app

import (
	"encoding/json"
	post "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	"github.com/salazarhugo/cheers1/services/notification-service/internal/notifications"
	"github.com/salazarhugo/cheers1/services/notification-service/internal/repository"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"net/http"
)

func PostSub(w http.ResponseWriter, r *http.Request) {
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

	event := &post.PostEvent{}
	err = proto.Unmarshal(m.Message.Data, event)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(event)

	users, err := repository.GetUsers([]string{event.UserId})
	if err != nil || len(users) < 1 {
		log.Println(err)
		return
	}
	user := users[0]

	repo := repository.NewRepository()
	postCreatorId := event.GetCreatorId()
	tokens, err := repo.GetUserTokens(postCreatorId)
	if err != nil {
		return
	}

	switch event.GetType() {
	case post.PostEvent_LIKE:
		data := notifications.LikePostNotification(user.Username, user.Picture)
		err := repo.SendNotification(map[string][]string{postCreatorId: tokens}, data)
		if err != nil {
			return
		}
	}
}