package repository

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"log"
)

func (r repository) ListComment(
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

		// Get the number of replies
		replyCount, err := r.redis.ZCard(ctx, getKeyReplyList(commentID)).Result()
		if err != nil {
			continue
		}

		// Get the number of likes
		likeCount, err := r.redis.SCard(ctx, getKeyCommentLikes(commentID)).Result()
		if err != nil {
			log.Println(err)
		}

		comments = append(comments, &comment.CommentItem{
			Comment:    com,
			UserItem:   userItem,
			ReplyCount: replyCount,
			LikeCount:  likeCount,
		})

	}

	return comments, nil
}
