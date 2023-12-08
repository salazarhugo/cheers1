package repository

import (
	uuid "github.com/google/uuid"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"log"
)

func (p *postRepository) CreatePost(
	userID string,
	post *Post,
) (string, error) {
	db := p.spanner

	post.ID = uuid.NewString()
	post.UserID = userID

	result := db.Create(&post)
	if result.Error != nil {
		return "", result.Error
	}

	go func() {
		err := pubsub.PublishProtoWithBinaryEncoding("post-topic", &pb.PostEvent{
			Event: &pb.PostEvent_Create{
				Create: &pb.CreatePost{
					Post:                      post.ToPostPb(),
					User:                      nil,
					SendNotificationToFriends: true,
				},
			},
		})
		if err != nil {
			log.Println(err)
		}
	}()

	return post.ID, nil
}
