package events

import (
	post "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"github.com/salazarhugo/cheers1/services/notification-service/internal/notifications"
	"github.com/salazarhugo/cheers1/services/notification-service/internal/repository"
	"github.com/salazarhugo/cheers1/services/notification-service/internal/service"
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
		user, err := service.GetUser(event.Like.UserId)
		if err != nil {
			log.Println(err)
			return
		}
		post, err := service.GetPost(event.Like.GetPostId())
		postCreatorId := post.CreatorId
		tokens, err := repo.GetUserTokens(postCreatorId)
		if err != nil {
			return
		}

		data := notifications.LikePostNotification(user.Username, user.Picture)
		err = repo.SendNotification(map[string][]string{postCreatorId: tokens}, data)
	case *post.PostEvent_Create:
		post := event.Create.Post
		creatorId := post.CreatorId
		user, err := service.GetUser(event.Create.GetPost().GetCreatorId())
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
			user.Name,
			user.Username,
			user.Picture,
			post.LocationName,
			post.Drink,
			post.Caption,
			len(post.Photos) > 0,
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
