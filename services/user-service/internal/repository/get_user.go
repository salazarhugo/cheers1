package repository

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/utils/models"
)

func (p *userRepository) GetUserWithViewer(
	viewerID string,
	otherUserID string,
) (*pb.GetUserResponse, error) {
	db := p.spanner

	var userItem models.UserWithViewer

	postCountQuery := db.
		Table("posts").
		Select("COUNT(*)").
		Where("user_id = users.UserId")

	friendCountQuery := db.
		Table("friendships").
		Select("COUNT(*)").
		Where("user_id2 = users.UserId")

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
		Table("users").
		Select(
			"users.*, EXISTS (?) as friend, EXISTS (?) AS requested, EXISTS (?) AS has_requested_viewer, (?) AS post_count , (?) AS friend_count",
			isFriendQuery,
			hasRequestedQuery,
			hasRequestedViewerQuery,
			postCountQuery,
			friendCountQuery,
		).
		Where("users.UserId = ? OR users.username = ?", otherUserID, otherUserID).
		Limit(1).
		First(&userItem).
		Error
	if err != nil {
		return nil, err
	}

	user := &pb.GetUserResponse{
		User:               userItem.ToUserPb(),
		PostCount:          userItem.PostCount,
		FriendsCount:       userItem.FriendCount,
		Friend:             userItem.Friend,
		Requested:          userItem.Requested,
		HasRequestedViewer: userItem.HasRequestedViewer,
	}

	return user, nil
}
