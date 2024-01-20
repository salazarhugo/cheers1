package repository

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/friendship-service/internal/model"
)

func (r repository) ListFriendRequests(
	page int,
	pageSize int,
	userId string,
) ([]*user.UserItem, error) {
	db := r.spanner

	limit, offset := utils.GetLimitAndOffsetPagination(page, pageSize)

	users := make([]*model.User, 0)

	err := db.
		Table("friend_requests").
		Select("users.*").
		Joins("JOIN users ON friend_requests.user_id1 = users.id").
		Where("user_id2 = ?", userId).
		Limit(limit).
		Offset(offset).
		Scan(&users).
		Error
	if err != nil {
		return nil, err
	}

	return model.ToUserItemsPb(users), nil
}
