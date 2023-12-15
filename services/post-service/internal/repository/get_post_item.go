package repository

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
)

func (p *postRepository) GetPostItem(
	viewerID string,
	postID string,
) (*pb.PostResponse, error) {
	db := p.spanner

	var post PostWithUserInfo
	result := db.Raw("SELECT posts.*, username, users.name, verified, picture, drinks.name, (select count(*) from post_likes where post_likes.post_id = posts.id) as likes, (SELECT EXISTS (SELECT 1 FROM post_likes WHERE user_id = ? AND post_id = posts.id)) as has_viewer_liked, FROM  posts JOIN users ON  posts.user_id = users.id LEFT OUTER JOIN drinks ON posts.drink_id = drinks.id WHERE posts.id = ? LIMIT ?", viewerID, postID, 1).Scan(&post)

	if result.Error != nil {
		return nil, result.Error
	}

	item := &pb.PostResponse{
		Post:         post.ToPostPb(),
		User:         post.ToUserPb(),
		LikeCount:    post.Likes,
		CommentCount: 0,
		HasLiked:     post.HasViewerLiked,
		IsCreator:    viewerID == post.UserID,
	}

	return item, nil
}
