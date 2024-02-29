package repository

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/note/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/mapper"
)

func (r *repository) FeedNote(
	userID string,
	page int,
	pageSize int,
) ([]*note.Note, error) {
	db := r.spanner
	limit, offset := utils.GetLimitAndOffsetPagination(page, pageSize)

	friendIDs, err := r.ListFriendIds(userID)
	if err != nil {
		return nil, err
	}

	// Add current user
	friendIDs = append(friendIDs, userID)

	notes := make([]*mapper.UserStatusItem, 0)

	result := db.
		Table("user_status").
		Select("user_status.*, users.*, drinks.DrinkId as drink_id, drinks.name as drink_name, drinks.icon as drink_icon").
		Joins("JOIN users ON user_status.UserId = users.UserId").
		Joins("LEFT OUTER JOIN drinks ON user_status.drink_id = drinks.DrinkId").
		Where("user_status.UserId IN ?", friendIDs).
		Where("TIMESTAMP_ADD(user_status.updated_at, INTERVAL 1 DAY) > CURRENT_TIMESTAMP()").
		Limit(limit).
		Offset(offset).
		Scan(&notes)

	if result.Error != nil {
		return nil, result.Error
	}

	items := make([]*note.Note, 0)
	for _, n := range notes {
		items = append(items, n.ToNotePb())
	}

	return items, nil
}
