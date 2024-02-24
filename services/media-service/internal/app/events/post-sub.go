package events

import (
	"github.com/labstack/gommon/log"
	post "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"github.com/salazarhugo/cheers1/services/media-service/internal/repository"
	"net/http"
)

func PostSub(w http.ResponseWriter, r *http.Request) {
	event := &post.PostEvent{}
	err := pubsub.UnmarshalPubSubMessage(r, event)
	if err != nil {
		log.Error(err)
		return
	}

	switch event := event.Event.(type) {
	case *post.PostEvent_Create:
		err = repository.HandlePostCreate(event.Create.GetPost().GetId())
	case *post.PostEvent_Delete:
		err = repository.HandlePostDelete(event.Delete.GetPost())
	default:
		log.Info("unhandled event")
	}

	if err != nil {
		log.Error(err)
		return
	}
}
