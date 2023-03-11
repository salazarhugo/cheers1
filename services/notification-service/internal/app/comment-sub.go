package app

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"github.com/salazarhugo/cheers1/services/notification-service/internal/notifications"
	"github.com/salazarhugo/cheers1/services/notification-service/internal/repository"
	"github.com/salazarhugo/cheers1/services/notification-service/internal/service"
	"log"
	"net/http"
)

func CommentSub(w http.ResponseWriter, r *http.Request) {
	event := &comment.CommentEvent{}
	err := pubsub.UnmarshalPubSubMessage(r, event)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println(event)
	repo := repository.NewRepository()

	switch event := event.Event.(type) {
	case *comment.CommentEvent_Created:
		user := event.Created.User
		comment := event.Created.Comment

		post, err := service.GetPost(comment.GetPostId())
		tokens, err := repo.GetUserTokens(post.CreatorId)
		if err != nil {
			return
		}

		data := notifications.CommentPostNotification(user.Username, user.Picture)
		err = repo.SendNotification(map[string][]string{post.CreatorId: tokens}, data)
	}

	if err != nil {
		log.Println(err)
	}

	return
}
