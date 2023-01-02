package app

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"github.com/salazarhugo/cheers1/services/post-service/internal/repository"
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
	repo := repository.NewPostRepository()

	switch event := event.Event.(type) {
	case *comment.CommentEvent_Created:
		err = repo.UpdatePostLastComment(event.Created.User, event.Created.Comment)
	case *comment.CommentEvent_Deleted:
	}

	if err != nil {
		http.Error(w, "failed to update post comments", http.StatusInternalServerError)
		return
	}

	return
}
