package repository

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
)

func (p *postRepository) ListPostLikes(
	postID string,
	page int,
	pageSize int,
) ([]*user.UserItem, error) {
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 18
	}
	skip := pageSize * (page - 1)

	var likes []LikeWithUser

	result := p.spanner.
		Table("post_likes").
		Select("post_likes.*, users.username, users.name, users.verified, users.picture").
		Joins("JOIN users ON post_likes.user_id = users.id").
		Where("post_id = ?", postID).
		Limit(pageSize).
		Offset(skip).
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
