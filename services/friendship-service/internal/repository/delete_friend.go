package repository

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/friendship/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
)

func (r repository) DeleteFriend(userId string, friendId string) error {
	err := r.redis.ZRem(context.Background(), getKeyFriends(userId), friendId).Err()
	err = r.redis.ZRem(context.Background(), getKeyFriends(friendId), userId).Err()

	if err != nil {
		return err
	}

	err = pubsub.PublishProtoWithBinaryEncoding("friendship-topic", &friendship.FriendshipEvent{
		Event: &friendship.FriendshipEvent_DeletedFriend{
			DeletedFriend: &friendship.DeletedFriend{
				From: userId,
				To:   friendId,
			},
		},
	})

	return err
}
