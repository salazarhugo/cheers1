package repository

import (
	"context"
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
		// Remove all replies
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

	return nil
}
