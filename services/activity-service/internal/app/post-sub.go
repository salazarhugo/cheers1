package app

import (
	"github.com/google/uuid"
	activity2 "github.com/salazarhugo/cheers1/gen/go/cheers/activity/v1"
	post "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"github.com/salazarhugo/cheers1/services/activity-service/internal/repository"
	"log"
	"net/http"
	"time"
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
		user := event.Like.User
		activity := &activity2.Activity{
			Id:           uuid.New().String(),
			Type:         activity2.Activity_LIKE_POST,
			Text:         user.Username + " liked your post",
			Picture:      user.Picture,
			UserId:       user.Id,
			Timestamp:    time.Now().Unix(),
			MediaPicture: "",
			MediaId:      "",
		}
		err = repo.CreateActivity(event.Like.Post.CreatorId, activity)
	}
}
