package repository

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/note/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/mapper"
	"github.com/salazarhugo/cheers1/services/note-service/internal/service"
)

func (r *repository) FeedNote(
	userID string,
	page int,
	pageSize int,
) ([]*note.Note, error) {
	db := r.spanner
	limit, offset := utils.GetLimitAndOffsetPagination(page, pageSize)

	friends, err := service.ListFriends(userID)
	if err != nil {
		return nil, err
	}

	// list of friend IDs
	friendIDs := make([]string, len(friends))
	for i, friend := range friends {
		friendIDs[i] = friend.Id
	}

	// Add current user
	friendIDs = append(friendIDs, userID)

	notes := make([]*mapper.UserStatusItem, 0)

	result := db.
		Table("user_status").
		Select("*").
		Joins("JOIN users ON user_status.UserId = users.UserId").
		Where("user_status.UserId IN ?", friendIDs).
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
