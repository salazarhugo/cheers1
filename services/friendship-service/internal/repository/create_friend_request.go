package repository

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/friendship/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
)

func (r repository) CreateFriendRequest(
	userId string,
	friendId string,
) error {
	db := r.spanner

	friendRequest := FriendRequest{
		UserId1: userId,
		UserId2: friendId,
	}

	err := db.Create(&friendRequest).Error
	if err != nil {
		return err
	}

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
