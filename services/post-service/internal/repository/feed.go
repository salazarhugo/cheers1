package repository

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
)

func (p *postRepository) FeedPost(
	userIDs []string,
	page int,
	pageSize int,
) (*pb.FeedPostResponse, error) {
	viewerID := userIDs[len(userIDs)-1]
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 18
	}

	skip := pageSize * (page - 1)

	var posts []PostWithUserInfo

	db := p.spanner
	result := db.Raw("SELECT posts.*, username, users.name, verified, picture, drinks.name, (select count(*) from post_likes where post_likes.post_id = posts.id) as likes, (SELECT EXISTS (SELECT 1 FROM post_likes WHERE user_id = ? AND post_id = posts.id)) as has_viewer_liked, FROM  posts JOIN users ON  posts.user_id = users.id LEFT OUTER JOIN drinks ON posts.drink_id = drinks.id WHERE posts.user_id IN ? LIMIT ? OFFSET ?", viewerID, userIDs, pageSize, skip).Scan(&posts)

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
