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
	viewerID := userIDs[len(userIDs)-1]
	limit, offset := utils.GetLimitAndOffsetPagination(page, pageSize)

	var posts []PostWithUserInfo

	db := p.spanner
	result := db.
		Table("posts").
		Select("posts.*, drinks.id as drink_id, drinks.name as drink_name, drinks.icon as drink_icon, users.username, users.name, users.verified, users.picture, (select count(*) from post_likes where post_likes.post_id = posts.id) as likes, (SELECT EXISTS (SELECT 1 FROM post_likes WHERE user_id = ? AND post_id = posts.id)) as has_viewer_liked", viewerID).
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
			LikeCount:    post.Likes,
			CommentCount: 0,
			HasLiked:     post.HasViewerLiked,
			IsCreator:    viewerID == post.UserID,
		}
		items = append(items, item)
	}

	nextPageToken := "postToken123"
	if len(posts) > 0 {
		nextPageToken = posts[len(posts)-1].ID
	}

	return &pb.FeedPostResponse{
		Posts:         items,
		NextPageToken: nextPageToken,
	}, nil
}
