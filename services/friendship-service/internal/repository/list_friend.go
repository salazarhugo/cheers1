package repository

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/models"
)

func (r repository) ListFriend(
	viewerID string,
	page int,
	pageSize int,
	userId string,
) ([]*user.UserItem, error) {
	db := r.spanner

	limit, offset := utils.GetLimitAndOffsetPagination(page, pageSize)

	users := make([]*models.UserWithViewer, 0)

	isFriendQuery := db.
		Table("friendships").
		Select("1").
		Where("user_id1 = users.UserId AND user_id2 = ?", viewerID)

	hasRequestedQuery := db.
		Table("friend_requests").
		Select("1").
		Where("user_id1 = ? AND user_id2 = users.UserId", viewerID)

	hasRequestedViewerQuery := db.
		Table("friend_requests").
		Select("1").
		Where("user_id1 = users.UserId AND user_id2 = ?", viewerID)

	err := db.
		Table("friendships").
		Select(
			"users.*, EXISTS (?) as friend, EXISTS (?) AS requested, EXISTS (?) AS has_requested_viewer",
			isFriendQuery,
			hasRequestedQuery,
			hasRequestedViewerQuery,
		).
		Joins("JOIN users ON friendships.user_id1 = users.UserId").
		Where("user_id2 = ?", userId).
		Limit(limit).
		Offset(offset).
		Scan(&users).
		Error
	if err != nil {
		return nil, err
	}

	return models.ToUserItemsPb2(users), nil
}
