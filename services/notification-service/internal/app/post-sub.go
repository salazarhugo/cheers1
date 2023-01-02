package app

import (
	post "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"github.com/salazarhugo/cheers1/services/notification-service/internal/notifications"
	"github.com/salazarhugo/cheers1/services/notification-service/internal/repository"
	"log"
	"net/http"
)

func PostSub(w http.ResponseWriter, r *http.Request) {
	event := &post.PostEvent{}
	err := pubsub.UnmarshalPubSubMessage(r, event)
	if err != nil {
		return
	}

	log.Println(event)

	repo := repository.NewRepository()

	switch event := event.Event.(type) {
	case *post.PostEvent_Like:
		users, err := repository.GetUsers([]string{event.Like.User.Id})
		if err != nil || len(users) < 1 {
			log.Println(err)
			return
		}
		user := users[0]
		postCreatorId := event.Like.Post.CreatorId
		tokens, err := repo.GetUserTokens(postCreatorId)
		if err != nil {
			return
		}

		data := notifications.LikePostNotification(user.Username, user.Picture)
		err = repo.SendNotification(map[string][]string{postCreatorId: tokens}, data)
		if err != nil {
			return
		}
	case *post.PostEvent_Create:
	case *post.PostEvent_Delete:
	default:
	}

}
