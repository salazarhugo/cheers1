package repository

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"log"
)

func (p *postRepository) DeletePostById(
	postID string,
) error {
	db := p.spanner
	result := db.Delete(&Post{PostId: postID})

	if result.Error != nil {
		return result.Error
	}

	go func() {
		err := pubsub.PublishProtoWithBinaryEncoding("post-topic", &pb.PostEvent{
			Event: &pb.PostEvent_Delete{
				Delete: &pb.DeletePost{
					Sender: nil,
				},
			},
		})
		if err != nil {
			log.Println(err)
		}
	}()

	return nil
}
