package repository

import (
	"context"
)

func (r repository) UnLikeComment(
	userID string,
	commentID string,
) error {
	ctx := context.Background()

	err := r.redis.SRem(
		ctx,
		getKeyCommentLikes(commentID),
		userID,
	).Err()
	if err != nil {
		return err
	}

	return nil
}
