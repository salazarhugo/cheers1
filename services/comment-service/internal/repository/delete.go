package repository

import (
	"context"
)

func (r repository) DeleteComment(
	postId string,
	commentId string,
) error {
	return r.redis.ZRem(context.Background(), getKeyPostComment(postId), commentId).Err()
}
