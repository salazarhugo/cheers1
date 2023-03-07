package app

import (
	"github.com/google/uuid"
	activity2 "github.com/salazarhugo/cheers1/gen/go/cheers/activity/v1"
	post "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"github.com/salazarhugo/cheers1/services/activity-service/internal/repository"
	"github.com/salazarhugo/cheers1/services/activity-service/internal/service"
	"log"
	"net/http"
	"time"
)

func PostSub(w http.ResponseWriter, r *http.Request) {
	event := &post.PostEvent{}
	err := pubsub.UnmarshalPubSubMessage(r, event)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(event)

	repo := repository.NewRepository()

	switch event := event.Event.(type) {
	case *post.PostEvent_Like:
		userId := event.Like.UserId
		postId := event.Like.PostId
		user, err := service.GetUser(userId)
		if err != nil {
			return
		}
		post, err := service.GetPost(postId)
		if err != nil {
			return
		}

		postThumbnail := ""
		if len(post.Photos) > 0 {
			postThumbnail = post.Photos[0]
		}

		activity := &activity2.Activity{
			Id:           uuid.New().String(),
			Type:         activity2.Activity_POST_LIKED,
			Text:         user.Username + " liked your post",
			Picture:      user.Picture,
			UserId:       user.Id,
			Timestamp:    time.Now().Unix(),
			MediaPicture: postThumbnail,
			MediaId:      post.Id,
			Username:     user.Username,
		}
		err = repo.CreateActivity(post.CreatorId, activity)
	}
}
