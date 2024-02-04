package repository

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
)

func (p *postRepository) GetPostItem(
	viewerID string,
	postID string,
) (*pb.PostResponse, error) {
	db := p.spanner
	var post PostItem

	mediaQuery := db.
		Raw("SELECT TO_JSON(t) FROM post_media AS t WHERE posts.PostId = t.PostId")

	likeCountQuery := db.
		Table("post_likes").
		Select("COUNT(*)").
		Where("post_likes.post_id = posts.PostId")

	hasViewerLikedQuery := db.
		Table("post_likes").
		Select("1").
		Where("user_id = ? AND post_id = posts.PostId", viewerID)

	result := db.
		Table("posts").
		Select(
			"posts.*, ARRAY(?) AS medias, username, users.name, verified, picture, drinks.id as drink_id, drinks.name as drink_name, drinks.icon as drink_icon, (?) AS like_count, EXISTS (?) AS has_viewer_liked",
			mediaQuery,
			likeCountQuery,
			hasViewerLikedQuery,
		).
		Joins("JOIN users ON posts.user_id = users.id").
		Joins("LEFT OUTER JOIN drinks ON posts.drink_id = drinks.id").
		Where("posts.PostId = ?", postID).
		First(&post)

	if result.Error != nil {
		return nil, result.Error
	}

	item := &pb.PostResponse{
		Post:         post.ToPostPb(),
		User:         post.ToUserPb(),
		LikeCount:    post.LikeCount,
		CommentCount: 0,
		HasLiked:     post.HasViewerLiked,
		IsCreator:    viewerID == post.UserID,
	}

	return item, nil
}
