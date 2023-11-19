package repository

import (
	"context"
	"fmt"
	"github.com/salazarhugo/cheers1/gen/go/cheers/friendship/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
)

func getKeyFriendRequests(userId string) string {
	return fmt.Sprintf("%s:%s", keyFriendRequests, userId)
}

func getKeyFriends(userId string) string {
	return fmt.Sprintf("%s:%s", keyFriends, userId)
}

func (r repository) CreateFriendRequest(
	userId string,
	friendId string,
) error {
	err := r.redis.SAdd(
		context.Background(),
		getKeyFriendRequests(friendId),
		userId,
	).Err()

	err = pubsub.PublishProtoWithBinaryEncoding("friendship-topic", &friendship.FriendshipEvent{
		Event: &friendship.FriendshipEvent_CreatedFriendRequest{
			CreatedFriendRequest: &friendship.CreatedFriendRequest{
				From: userId,
				To:   friendId,
			},
		},
	})

	return err
}
