package repository

import (
	"gorm.io/gorm"
)

func (r repository) CreateFriend(userId string, friendId string) error {
	db := r.spanner

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&Friendship{
			UserId1: userId,
			UserId2: friendId,
		}).Error; err != nil {
			return err
		}

		// If self-friend then just add one record
		if userId == friendId {
			return nil
		}

		if err := tx.Create(&Friendship{
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

	return nil
}
