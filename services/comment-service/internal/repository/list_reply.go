package repository

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
)

func (repository repository) ListReplies(
	userID string,
	parentCommentId string,
) ([]*comment.CommentItem, error) {
	ctx := context.Background()
	// Get the IDs of all comments for the post
	commentIDs, err := repository.redis.ZRange(ctx, getKeyReplyList(parentCommentId), 0, -1).Result()
	if err != nil {
		return nil, err
	}

	var comments []*comment.CommentItem
	for _, commentID := range commentIDs {
		item, err := repository.GetCommentItem(userID, commentID)
		if err != nil {
			continue
		}

		comments = append(comments, item)
	}

	return comments, nil
}
