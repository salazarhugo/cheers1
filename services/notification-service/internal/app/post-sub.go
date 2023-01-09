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
		user, err := repository.GetUser(event.Like.User.Id)
		if err != nil {
			log.Println(err)
			return
		}
		postCreatorId := event.Like.Post.CreatorId
		tokens, err := repo.GetUserTokens(postCreatorId)
		if err != nil {
			return
		}

		data := notifications.LikePostNotification(user.Username, user.Picture)
		err = repo.SendNotification(map[string][]string{postCreatorId: tokens}, data)
	case *post.PostEvent_Create:
		creatorId := event.Create.Post.CreatorId
		user, err := repository.GetUser(event.Create.GetPost().GetCreatorId())
		if err != nil {
			log.Println(err)
			return
		}

		friends, err := repository.ListFriend(creatorId)
		if err != nil {
			log.Println(err)
			return
		}

		usersWithTokens, err := repo.GetUsersTokens(friends)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(usersWithTokens)

		data := notifications.CreatePostNotification(
			user.Username,
			user.Picture,
			event.Create.Post.Drink,
		)

		err = repo.SendNotification(usersWithTokens, data)
	case *post.PostEvent_Delete:
	default:
	}

	if err != nil {
		log.Println(err)
		return
	}
}
