package repository

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/friendship-service/internal/model"
)

func (r repository) ListFriend(
	viewerID string,
	page int,
	pageSize int,
	userId string,
) ([]*user.UserItem, error) {
	db := r.spanner

	limit, offset := utils.GetLimitAndOffsetPagination(page, pageSize)

	users := make([]*model.UserItem, 0)

	err := db.
		Table("friendships").
		Select("users.*, EXISTS (SELECT 1 FROM friendships WHERE user_id1 = id AND user_id2 = ?) as friend, EXISTS (SELECT 1 FROM friend_requests WHERE user_id1 = ? AND user_id2 = users.id) AS requested, EXISTS (SELECT 1 FROM friend_requests WHERE user_id1 = users.id AND user_id2 = ?) AS has_requested_viewer", viewerID, viewerID, viewerID).
		Joins("JOIN users ON friendships.user_id1 = users.id").
		Where("user_id2 = ?", userId).
		Limit(limit).
		Offset(offset).
		Scan(&users).
		Error
	if err != nil {
		return nil, err
	}

	return model.ToUserItemsPb2(users), nil
}
