package repository

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
)

func (r repository) GetCommentItem(
	commentID string,
) (*comment.CommentItem, error) {
	ctx := context.Background()

	com := &comment.Comment{}
	commentFields, err := r.redis.HGetAll(ctx, getKeyComment(commentID)).Result()
	if err != nil {
		return nil, err
	}

	err = utils.MapToProto(com, commentFields)
	if err != nil {
		return nil, err
	}

	userItem, err := r.GetUserItem(com.UserId)
	if err != nil {
		return nil, err
	}

	// Get the number of replies
	replyCount, err := r.redis.ZCard(ctx, getKeyReplyList(commentID)).Result()
	if err != nil {
		return nil, err
	}

	item := &comment.CommentItem{
		Comment:    com,
		UserItem:   userItem,
		ReplyCount: replyCount,
	}

	return item, nil
}
