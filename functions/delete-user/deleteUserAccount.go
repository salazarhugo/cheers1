package delete_user

import (
	"context"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"time"
)

func init() {
}

// AuthEvent is the payload of a Firestore Auth event.
type AuthEvent struct {
	Email    string `json:"email"`
	Metadata struct {
		CreatedAt time.Time `json:"createdAt"`
	} `json:"metadata"`
	UID string `json:"uid"`
}

// DeleteAuth is triggered by Firestore Auth Delete event.
func DeleteAuth(ctx context.Context, e AuthEvent) error {

	err := pubsub.PublishProtoWithBinaryEncoding("user-topic", &pb.UserEvent{
		Event: &pb.UserEvent_Create{
			Create: &pb.CreateUser{
				User: response.User,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return nil
}
