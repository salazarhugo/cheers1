package repository

import (
	"context"
	"strconv"
)

func (p *postRepository) GetCommentCount(postID string) (int, error) {
	ctx := context.Background()

	// Get the comment count for the post from Redis
	countStr, err := p.redis.HGet(ctx, getKeyPostComments(), getKeyPost(postID)).Result()
	if err != nil {
		return 0, err
	}

	// Convert the count to an integer
	count, err := strconv.Atoi(countStr)
	if err != nil {
		return 0, err
	}

	return count, nil
}