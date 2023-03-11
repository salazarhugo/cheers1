package repository

import (
	"context"
	"fmt"
)

const (
	keyUser     = "user"
	keyPost     = "post"
	keyComment  = "comment"
	keyComments = "comments"
	keyReplies  = "replies"
)

func getKeyPostComments() string {
	return fmt.Sprintf("%s:%s", keyPost, keyComments)
}

func getKeyPost(postId string) string {
	return fmt.Sprintf("%s:%s", keyPost, postId)
}

func (p *postRepository) IncrementCommentCount(
	postID string,
) error {
	ctx := context.Background()

	// Increment the comment count of the post
	err := p.redis.HIncrBy(ctx, getKeyPostComments(), getKeyPost(postID), 1).Err()
	if err != nil {
		return err
	}

	return nil
}
