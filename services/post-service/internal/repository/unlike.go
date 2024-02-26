package repository

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	"github.com/salazarhugo/cheers1/libs/utils/models"
)

func (p *postRepository) UnlikePost(
	userID string,
	postID string,
) (*pb.UnlikePostResponse, error) {
	db := p.spanner

	like := &models.Like{
		UserID: userID,
		PostID: postID,
	}
	result := db.Table("post_likes").Delete(like)
	if result.Error != nil {
		return nil, result.Error
	}

	return &pb.UnlikePostResponse{
		Success: true,
	}, nil
}
