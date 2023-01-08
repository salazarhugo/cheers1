package repository

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/friendship/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r repository) AcceptFriendRequest(userId string, friendId string) error {
	ctx := context.Background()

	exists, err := r.redis.Exists(ctx, getKeyFriendRequests(userId), friendId).Result()
	if exists < 1 {
		return status.Error(codes.NotFound, "Friend request not found")
	}

	// Delete friend request
	err = r.redis.SRem(ctx, getKeyFriendRequests(userId), friendId).Err()
	if err != nil {
		return err
	}

	err = r.redis.SRem(ctx, getKeyFriendRequests(friendId), userId).Err()
	if err != nil {
		return err
	}

	// Add friend to my friend list
	err = r.redis.SAdd(
		context.Background(),
		getKeyFriends(userId),
		friendId,
	).Err()
	if err != nil {
		return err
	}

	// Add me to other user friend list
	err = r.redis.SAdd(
		context.Background(),
		getKeyFriends(friendId),
		userId,
	).Err()
	if err != nil {
		return err
	}

	err = pubsub.PublishProtoWithBinaryEncoding("friendship-topic", &friendship.FriendshipEvent{
		Event: &friendship.FriendshipEvent_CreatedFriend{
			CreatedFriend: &friendship.CreatedFriend{
				From: userId,
				To:   friendId,
			},
		},
	})

	return err
}
