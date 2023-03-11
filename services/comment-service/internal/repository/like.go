package repository

import (
	"context"
	"fmt"
	commentpb "github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"log"
)

func getKeyCommentLikes(commentID string) string {
	return fmt.Sprintf("%s:%s:%s", keyComment, commentID, keyLikes)
}

func (r repository) LikeComment(
	userID string,
	commentID string,
) error {
	ctx := context.Background()

	err := r.redis.SAdd(
		ctx,
		getKeyCommentLikes(commentID),
		userID,
	).Err()
	if err != nil {
		return err
	}

	go func() {
		err := pubsub.PublishProtoWithBinaryEncoding("comment-topic", &commentpb.CommentEvent{
			Event: &commentpb.CommentEvent_Liked{
				Liked: &commentpb.LikedComment{
					CommentId: commentID,
					UserId:    userID,
				},
			},
		})
		if err != nil {
			log.Println(err)
		}
	}()

	return nil
}
