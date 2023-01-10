package repository

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/salazarhugo/cheers1/gen/go/cheers/friendship/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"log"
)

func (r repository) DeleteFriend(userId string, friendId string) error {
	ctx := context.Background()
	err := r.redis.Watch(ctx, func(tx *redis.Tx) error {
		err := tx.SRem(ctx, getKeyFriends(userId), friendId).Err()
		err = tx.SRem(ctx, getKeyFriends(friendId), userId).Err()
		return err
	}, "")

	if err != nil {
		log.Println(err)
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
