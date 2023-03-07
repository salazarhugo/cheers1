package repository

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	commentpb "github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"github.com/salazarhugo/cheers1/services/comment-service/internal/domain"
	"log"
	"time"
)

func getKeyReplyList(commentId string) string {
	return fmt.Sprintf("%s:%s:%s", keyComment, commentId, keyReplies)
}

func (r repository) CreateReplyComment(
	userId string,
	text string,
	postId string,
	replyCommentId string,
) error {
	ctx := context.Background()
	replyComment := &commentpb.Comment{
		Id:         uuid.New().String(),
		Text:       text,
		CreateTime: time.Now().Unix(),
		UserId:     userId,
		PostId:     postId,
	}
	mentions := domain.GetMentions(text)

	m, err := utils.ProtoToMap(replyComment)
	if err != nil {
		return err
	}

	// Store the reply replyComment details in a hash
	err = r.redis.HSet(
		ctx,
		getKeyComment(replyComment.Id),
		m,
	).Err()
	if err != nil {
		return err
	}

	// Add the reply replyComment ID to the replyComment's reply list
	err = r.redis.ZAdd(
		ctx,
		getKeyReplyList(postId),
		redis.Z{
			Score:  float64(time.Now().Unix()),
			Member: replyComment.Id,
		},
	).Err()
	if err != nil {
		return err
	}

	userItem, err := r.GetUserItem(replyComment.UserId)

	item := &commentpb.CommentItem{
		Comment:  replyComment,
		UserItem: userItem,
	}

	go func() {
		err := pubsub.PublishProtoWithBinaryEncoding("comment-topic", &commentpb.CommentEvent{
			Event: &commentpb.CommentEvent_Created{
				Created: &commentpb.CreatedComment{
					Comment:  item.Comment,
					User:     item.UserItem,
					Mentions: mentions,
				},
			},
		})
		if err != nil {
			log.Println(err)
		}
	}()

	return nil
}
