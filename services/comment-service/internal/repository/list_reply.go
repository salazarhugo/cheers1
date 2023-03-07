package repository

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
)

func (r repository) ListReplies(
	parentCommentId string,
) ([]*comment.CommentItem, error) {
	ctx := context.Background()
	// Get the IDs of all comments for the post
	commentIDs, err := r.redis.ZRange(ctx, getKeyReplyList(parentCommentId), 0, -1).Result()
	if err != nil {
		return nil, err
	}

	// Get the details of all comments and their first three replies
	var comments []*comment.CommentItem
	for _, commentID := range commentIDs {
		commentFields, err := r.redis.HGetAll(ctx, getKeyComment(commentID)).Result()
		if err != nil {
			continue
		}
		com := &comment.Comment{}
		err = utils.MapToProto(com, commentFields)

		userItem, err := r.GetUserItem(com.UserId)
		if err != nil {
			continue
		}

		comments = append(comments, &comment.CommentItem{
			Comment:  com,
			UserItem: userItem,
		})
	}

	return comments, nil
}
