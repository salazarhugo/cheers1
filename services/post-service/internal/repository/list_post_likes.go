package repository

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"github.com/salazarhugo/cheers1/libs/utils"
)

func (p *postRepository) ListPostLikes(
	viewerID string,
	postID string,
	page int,
	pageSize int,
) ([]*user.UserItem, error) {
	limit, offset := utils.GetLimitAndOffsetPagination(page, pageSize)

	var likes []LikeWithUser

	result := p.spanner.
		Table("post_likes").
		Select("post_likes.*, users.*, EXISTS (SELECT 1 FROM friendships WHERE user_id1 = users.id AND user_id2 = ?) as friend, EXISTS (SELECT 1 FROM friend_requests WHERE user_id1 = ? AND user_id2 = users.id) AS requested, EXISTS (SELECT 1 FROM friend_requests WHERE user_id1 = users.id AND user_id2 = ?) AS has_requested_viewer", viewerID, viewerID, viewerID).
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
			Id:                 l.UserID,
			Name:               l.Name,
			Username:           l.Username,
			Verified:           l.Verified,
			Picture:            l.Picture,
			HasFollowed:        false,
			StoryState:         0,
			Friend:             l.Friend,
			Requested:          l.Requested,
			HasRequestedViewer: l.HasRequestedViewer,
			Banner:             "",
		}
		users = append(users, u)
	}

	return users, nil
}
