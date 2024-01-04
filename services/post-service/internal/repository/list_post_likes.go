package repository

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"github.com/salazarhugo/cheers1/libs/utils"
)

func (p *postRepository) ListPostLikes(
	postID string,
	page int,
	pageSize int,
) ([]*user.UserItem, error) {
	limit, offset := utils.GetLimitAndOffsetPagination(page, pageSize)

	var likes []LikeWithUser

	result := p.spanner.
		Table("post_likes").
		Select("post_likes.*, users.username, users.name, users.verified, users.picture").
		Joins("JOIN users ON post_likes.user_id = users.id").
		Where("post_id = ?", postID).
		Limit(limit).
		Offset(offset).
		Scan(&likes)

	if result.Error != nil {
		return nil, result.Error
	}

	users := make([]*user.UserItem, 0)

	for _, l := range likes {
		u := &user.UserItem{
			Id:       l.UserID,
			Name:     l.Name,
			Username: l.Username,
			Verified: l.Verified,
			Picture:  l.Picture,
		}
		users = append(users, u)
	}

	return users, nil
}
