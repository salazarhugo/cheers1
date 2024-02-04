package repository

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
)

func (p *postRepository) FeedPost(
	userIDs []string,
	page int,
	pageSize int,
) (*pb.FeedPostResponse, error) {
	db := p.spanner
	viewerID := userIDs[len(userIDs)-1]
	limit, offset := utils.GetLimitAndOffsetPagination(page, pageSize)
	var posts []PostItem

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
			"posts.*, ARRAY(?) AS medias, drinks.id as drink_id, drinks.name as drink_name, drinks.icon as drink_icon, users.username, users.name, users.verified, users.picture, (?) as like_count, (?) as has_viewer_liked",
			mediaQuery,
			likeCountQuery,
			hasViewerLikedQuery,
		).
		Joins("JOIN users ON  posts.user_id = users.id").
		Joins("LEFT OUTER JOIN drinks ON posts.drink_id = drinks.id").
		Where("posts.user_id IN ?", userIDs).
		Order("posts.created_at DESC").
		Limit(limit).
		Offset(offset).
		Scan(&posts)

	if result.Error != nil {
		return nil, result.Error
	}

	items := make([]*pb.PostResponse, 0)
	for _, post := range posts {
		item := &pb.PostResponse{
			Post:         post.ToPostPb(),
			User:         post.ToUserPb(),
			LikeCount:    post.LikeCount,
			CommentCount: 0,
			HasLiked:     post.HasViewerLiked,
			IsCreator:    viewerID == post.UserID,
		}
		items = append(items, item)
	}

	nextPageToken := "postToken123"
	if len(posts) > 0 {
		nextPageToken = posts[len(posts)-1].PostId
	}

	return &pb.FeedPostResponse{
		Posts:         items,
		NextPageToken: nextPageToken,
	}, nil
}
