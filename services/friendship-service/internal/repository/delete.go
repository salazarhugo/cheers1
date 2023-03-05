package repository

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/friendship/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
)

func (r repository) DeleteFriendRequest(
	userId string,
	friendId string,
) error {
	err := r.redis.SRem(context.Background(), getKeyFriendRequests(userId), friendId).Err()
	if err != nil {
		return err
	}

	err = pubsub.PublishProtoWithBinaryEncoding("friendship-topic", &friendship.FriendshipEvent{
		Event: &friendship.FriendshipEvent_DeletedFriendRequest{
			DeletedFriendRequest: &friendship.DeletedFriendRequest{
				From: userId,
				To:   friendId,
			},
		},
	})

	return err
}
