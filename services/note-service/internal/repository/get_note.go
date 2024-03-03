package repository

import (
	"github.com/salazarhugo/cheers1/libs/utils/mapper"
)

func (r *repository) GetNoteItem(
	userID string,
) (*mapper.UserStatusItem, error) {
	db := r.spanner

	var userStatus mapper.UserStatusItem

	result := db.
		Table("user_status").
		Select("user_status.*, users.*, drinks.DrinkId as drink_id, drinks.name as drink_name, drinks.icon as drink_icon").
		Joins("JOIN users ON user_status.UserId = users.UserId").
		Joins("LEFT OUTER JOIN drinks ON user_status.drink_id = drinks.DrinkId").
		Where("user_status.UserId = ?", userID).
		Where("TIMESTAMP_ADD(user_status.updated_at, INTERVAL 1 DAY) > CURRENT_TIMESTAMP()").
		First(&userStatus)

	if result.Error != nil {
		return nil, result.Error
	}

	return &userStatus, nil
}
