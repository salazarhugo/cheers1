package app

import (
	"fmt"
	"github.com/google/uuid"
	activity2 "github.com/salazarhugo/cheers1/gen/go/cheers/activity/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"github.com/salazarhugo/cheers1/services/activity-service/internal/repository"
	"github.com/salazarhugo/cheers1/services/activity-service/internal/service"
	"log"
	"net/http"
	"time"
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
		mentions := event.Created.Mentions
		comment := event.Created.Comment
		commentCreator, err := service.GetUser(event.Created.User.Id)
		if err != nil {
			return
		}
		post, err := service.GetPost(comment.PostId)
		if err != nil {
			return
		}
		postThumbnail := ""
		if len(post.Photos) > 0 {
			postThumbnail = post.Photos[0]
		}

		activityMention := &activity2.Activity{
			Id:           uuid.New().String(),
			Type:         activity2.Activity_MENTION_POST_COMMENT,
			Text:         fmt.Sprintf("%s mentioned you in a comment: %s", commentCreator.Username, comment.Text),
			Username:     commentCreator.Username,
			Picture:      commentCreator.Picture,
			UserId:       commentCreator.Id,
			Timestamp:    time.Now().Unix(),
			MediaPicture: postThumbnail,
			MediaId:      post.Id,
		}

		activityComment := &activity2.Activity{
			Id:           uuid.New().String(),
			Type:         activity2.Activity_POST_COMMENTED,
			Text:         fmt.Sprintf("%s commented: %s", commentCreator.Username, comment.Text),
			Username:     commentCreator.Username,
			Picture:      commentCreator.Picture,
			UserId:       commentCreator.Id,
			Timestamp:    time.Now().Unix(),
			MediaPicture: postThumbnail,
			MediaId:      post.Id,
		}
		err = repo.CreateActivity(post.CreatorId, activityComment)

		for _, mention := range mentions {
			userMentioned, err := service.GetUser(mention)
			if err != nil {
				continue
			}
			// Skip the creator of the post
			if userMentioned.Id == post.CreatorId {
				continue
			}
			err = repo.CreateActivity(userMentioned.Id, activityMention)
		}
	case *comment.CommentEvent_Deleted:
	}

	if err != nil {
		log.Println(err)
		return
	}

	return
}
