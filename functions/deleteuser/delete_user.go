package deleteuser

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/auth/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"time"
)

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
	err := pubsub.PublishProtoWithBinaryEncoding("auth-topic", &auth.AuthEvent{
		Event: &auth.AuthEvent_Deleted{
			Deleted: &auth.DeletedAuth{
				Email: e.Email,
				Uid:   e.UID,
			},
		},
	})

	if err != nil {
		return err
	}

	return nil
}
