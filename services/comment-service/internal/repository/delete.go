package repository

import (
	"context"
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

	comment.PostId

	// Delete comment hash
	err := r.redis.HDel(
		ctx,
		getKeyComment(commentId),
	).Err()

	err := r.redis.ZRem(
		ctx,
		getKeyPostComment(postId), commentId,
	).Err()

	return err
}
