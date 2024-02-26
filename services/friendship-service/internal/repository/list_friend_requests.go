package repository

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/models"
)

func (r repository) ListFriendRequests(
	page int,
	pageSize int,
	userId string,
) ([]*user.UserItem, error) {
	db := r.spanner

	limit, offset := utils.GetLimitAndOffsetPagination(page, pageSize)

	users := make([]*models.User, 0)

	err := db.
		Table("friend_requests").
		Select("users.*").
		Joins("JOIN users ON friend_requests.user_id1 = users.UserId").
		Where("user_id2 = ?", userId).
		Limit(limit).
		Offset(offset).
		Scan(&users).
		Error
	if err != nil {
		return nil, err
	}

	return models.ToUserItemsPb(users), nil
}
