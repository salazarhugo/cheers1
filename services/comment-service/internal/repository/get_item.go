package repository

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"log"
)

func (repository repository) GetCommentItem(
	viewerID string,
	commentID string,
) (*comment.CommentItem, error) {
	ctx := context.Background()

	com := &comment.Comment{}
	commentFields, err := repository.redis.HGetAll(ctx, getKeyComment(commentID)).Result()
	if err != nil {
		return nil, err
	}

	err = utils.MapToProto(com, commentFields)
	if err != nil {
		return nil, err
	}

	userItem, err := repository.GetUserItem(com.UserId)
	if err != nil {
		return nil, err
	}

	// Get the number of replies
	replyCount, err := repository.redis.ZCard(ctx, getKeyReplyList(commentID)).Result()
	if err != nil {
		return nil, err
	}

	// Get the number of likes
	likeCount, err := repository.redis.SCard(ctx, getKeyCommentLikes(commentID)).Result()
	if err != nil {
		log.Println(err)
	}

	// Check if user has liked the comment
	hasLiked, err := repository.HasUserLikedComment(viewerID, commentID)
	if err != nil {
		log.Println(err)
	}

	item := &comment.CommentItem{
		Comment:    com,
		UserItem:   userItem,
		ReplyCount: replyCount,
		LikeCount:  likeCount,
		HasLiked:   hasLiked,
	}

	return item, nil
}
