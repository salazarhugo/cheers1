package repository

import (
	"context"
	commentpb "github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"log"
)

func (r repository) DeleteComment(
	commentId string,
) error {
	ctx := context.Background()

	// Get comment
	comment, err := r.GetComment(commentId)
	if err != nil {
		return err
	}

	isReply := comment.ReplyToCommentId != ""

	pipeline := r.redis.TxPipeline()

	if isReply {
		// Remove the ID from the list of replies of the parent comment
		pipeline.ZRem(
			ctx,
			getKeyReplyList(comment.ReplyToCommentId),
			comment.Id,
		)
	} else {
		// Remove the ID from the list comments of that post
		pipeline.ZRem(
			ctx,
			getKeyPostComment(comment.PostId),
			commentId,
		)

		// Get all replies
		replyIDs, err := pipeline.ZRange(ctx, getKeyReplyList(comment.Id), 0, -1).Result()
		if err != nil {
		}

		// Delete all replies hash
		for _, replyID := range replyIDs {
			pipeline.Del(
				ctx,
				getKeyComment(replyID),
			)
		}

		// Remove reply list
		pipeline.Del(
			ctx,
			getKeyReplyList(comment.Id),
		)
	}

	// Delete comment hash
	pipeline.Del(
		ctx,
		getKeyComment(commentId),
	)

	_, err = pipeline.Exec(ctx)
	if err != nil {
		log.Println(err)
		return err
	}

	go func() {
		err := pubsub.PublishProtoWithBinaryEncoding("comment-topic", &commentpb.CommentEvent{
			Event: &commentpb.CommentEvent_Deleted{
				Deleted: &commentpb.DeletedComment{
					Comment: comment,
					User:    nil,
				},
			},
		})
		if err != nil {
			log.Println(err)
		}
	}()

	return nil
}
