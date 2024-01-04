package events

import (
	"github.com/labstack/gommon/log"
	post "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"github.com/salazarhugo/cheers1/services/notification-service/internal/notifications"
	"github.com/salazarhugo/cheers1/services/notification-service/internal/repository"
	"github.com/salazarhugo/cheers1/services/notification-service/internal/service"
	"net/http"
)

func PostSub(w http.ResponseWriter, r *http.Request) {
	event := &post.PostEvent{}
	err := pubsub.UnmarshalPubSubMessage(r, event)
	if err != nil {
		return
	}

	log.Info(event)

	repo := repository.NewRepository()

	switch event := event.Event.(type) {
	case *post.PostEvent_Like:
		user, err := service.GetUser(event.Like.UserId)
		if err != nil {
			log.Error(err)
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
		if event.Create.SendNotificationToFriends == false {
			log.Info("skipping: send notification to friends is false.")
			return
		}
		post := event.Create.Post
		creatorId := post.CreatorId
		user, err := service.GetUser(event.Create.GetPost().GetCreatorId())
		if err != nil {
			log.Error(err)
			return
		}

		friends, err := repository.ListFriend(creatorId)
		if err != nil {
			log.Error(err)
			return
		}

		usersWithTokens, err := repo.GetUsersTokens(friends)
		if err != nil {
			log.Error(err)
			return
		}
		log.Debug(usersWithTokens)

		data := notifications.CreatePostNotification(
			user.Name,
			user.Username,
			user.Picture,
			post.LocationName,
			post.Drink.Name,
			post.Caption,
			len(post.Photos) > 0,
		)

		err = repo.SendNotification(usersWithTokens, data)
	case *post.PostEvent_Delete:
	default:
	}

	if err != nil {
		log.Error(err)
		return
	}
}
