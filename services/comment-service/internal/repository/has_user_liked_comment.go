package repository

import (
	"context"
)

func (repository repository) HasUserLikedComment(
	userID string,
	commentID string,
) (bool, error) {
	ctx := context.Background()

	// Check if the user ID is a member of the set of likes associated with the comment
	isMember, err := repository.redis.SIsMember(ctx, getKeyCommentLikes(commentID), userID).Result()
	if err != nil {
		return false, err
	}

	return isMember, nil
}
