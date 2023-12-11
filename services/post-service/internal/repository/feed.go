package repository

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
)

func (p *postRepository) FeedPost(
	userIDs []string,
	page int,
	pageSize int,
) (*pb.FeedPostResponse, error) {
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 18
	}

	skip := pageSize * (page - 1)

	var posts []PostWithUserInfo

	db := p.spanner
	result := db.Table("posts").Select("posts.*, username, users.name, verified, picture, drinks.id as drink_id, drinks.name as drink_name, drinks.icon as drink_icon").Joins("JOIN users ON posts.user_id = users.id").Joins("LEFT OUTER JOIN drinks ON posts.drink_id = drinks.id").Where("user_id IN ?", userIDs).Limit(int(pageSize)).Offset(int(skip)).Order("created_at asc").Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}

	items := make([]*pb.PostResponse, 0)
	for _, post := range posts {
		item := &pb.PostResponse{
			Post:         post.ToPostPb(),
			User:         post.ToUserPb(),
			LikeCount:    0,
			CommentCount: 0,
			HasLiked:     false,
			IsCreator:    false,
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
