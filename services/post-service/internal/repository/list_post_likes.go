package repository

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/models"
)

func (p *postRepository) ListPostLikes(
	viewerID string,
	postID string,
	page int,
	pageSize int,
) ([]*user.UserItem, error) {
	db := p.spanner
	limit, offset := utils.GetLimitAndOffsetPagination(page, pageSize)

	var likes []models.LikeWithUser

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

	result := db.
		Table("post_likes").
		Select(
			"post_likes.*, users.*, EXISTS (?) as friend, EXISTS (?) AS requested, EXISTS (?) AS has_requested_viewer",
			isFriendQuery,
			hasRequestedQuery,
			hasRequestedViewerQuery,
		).
		Joins("JOIN users ON post_likes.user_id = users.UserId").
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
