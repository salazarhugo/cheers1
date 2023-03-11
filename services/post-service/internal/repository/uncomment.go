package repository

import (
	"context"
)

func (p *postRepository) DecrementCommentCount(
	postID string,
) error {
	ctx := context.Background()

	// Decrement the comment count of the post
	err := p.redis.HIncrBy(ctx, getKeyPostComments(), getKeyPost(postID), -1).Err()
	if err != nil {
		return err
	}

	return nil
}
