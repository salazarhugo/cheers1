package repository

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/friendship/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r repository) AcceptFriendRequest(
	userId string,
	friendId string,
) error {
	friendRequest, err := r.GetFriendRequest(friendId, userId)
	if err != nil {
		return status.Error(codes.NotFound, "Friend request not found")
	}

	// Delete friend request
	err = r.DeleteFriendRequest(friendRequest.UserId1, friendRequest.UserId2)
	if err != nil {
		return err
	}

	// Create friendship
	err = r.CreateFriend(userId, friendId)
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
