package repository

import (
	"context"
)

func (repository repository) UnLikeComment(
	userID string,
	commentID string,
) error {
	ctx := context.Background()

	err := repository.redis.SRem(
		ctx,
		getKeyCommentLikes(commentID),
		userID,
	).Err()
	if err != nil {
		return err
	}

	return nil
}
