package repository

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
)

func (r repository) ListReplyComment(
	parentCommentId string,
	postID string,
) ([]*comment.CommentItem, error) {
	ctx := context.Background()
	// Get the IDs of all comments for the post
	commentIDs, err := r.redis.ZRange(ctx, getKeyPostComment(postID), 0, -1).Result()
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

		// Get the IDs of the replies for the comment, ordered by their timestamp
		replyIDs, err := r.redis.ZRevRange(ctx, getKeyReplyList(commentID), 0, 2).Result()
		if err != nil {
			continue
		}

		// Get the details of the reply comments and add them to the comment object
		for _, replyID := range replyIDs {
			replyFields, err := r.redis.HGetAll(ctx, getKeyComment(replyID)).Result()
			if err != nil {
				continue
			}
			com := &comment.Comment{}
			err = utils.MapToProto(com, replyFields)

			userItem, err := r.GetUserItem(com.UserId)
			if err != nil {
				continue
			}

			comments = append(comments, &comment.CommentItem{
				Comment:  com,
				UserItem: userItem,
			})
		}
	}

	return comments, nil
}
