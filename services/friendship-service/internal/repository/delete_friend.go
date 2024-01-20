package repository

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/friendship/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"gorm.io/gorm"
)

func (r repository) DeleteFriend(userId string, friendId string) error {
	db := r.spanner

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&Friendship{
			UserId1: userId,
			UserId2: friendId,
		}).Error; err != nil {
			return err
		}

		if err := tx.Delete(&Friendship{
			UserId1: friendId,
			UserId2: userId,
		}).Error; err != nil {
			return err
		}

		return nil
	})
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
