package repository

import "C"
import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	"github.com/salazarhugo/cheers1/libs/utils/models"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"log"
)

func (p *postRepository) LikePost(
	userID string,
	postID string,
) (*pb.LikePostResponse, error) {
	db := p.spanner

	like := &models.Like{
		UserID: userID,
		PostID: postID,
	}

	result := db.
		Table("post_likes").
		Create(&like)
	if result.Error != nil {
		return nil, result.Error
	}

	go func() {
		err := pubsub.PublishProtoWithBinaryEncoding("post-topic", &pb.PostEvent{
			Event: &pb.PostEvent_Like{
				Like: &pb.LikePost{
					PostId: postID,
					UserId: userID,
				},
			},
		})
		if err != nil {
			log.Println(err)
		}
	}()

	return &pb.LikePostResponse{
		Success: true,
	}, nil
}
