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

func getKeyComment(commentId string) string {
	return fmt.Sprintf("%s:%s", keyComment, commentId)
}

func getKeyPostComment(postId string) string {
	return fmt.Sprintf("%s:%s:%s", keyPost, postId, keyComments)
}

func (r repository) CreateComment(
	userId string,
	text string,
	postId string,
) (string, error) {
	ctx := context.Background()
	commentID := uuid.New().String()
	comment := &commentpb.Comment{
		Id:         commentID,
		Text:       text,
		CreateTime: time.Now().Unix(),
		UserId:     userId,
		PostId:     postId,
	}
	mentions := domain.GetMentions(text)

	m, err := utils.ProtoToMap(comment)
	if err != nil {
		return "", err
	}

	// Store the comment details in a hash
	err = r.redis.HSet(
		ctx,
		getKeyComment(comment.Id),
		m,
	).Err()
	if err != nil {
		return "", err
	}

	// Add the comment ID to the post's comment set
	err = r.redis.ZAdd(
		ctx,
		getKeyPostComment(postId),
		redis.Z{
			Score:  float64(time.Now().Unix()),
			Member: comment.Id,
		},
	).Err()
	if err != nil {
		return "", err
	}

	userItem, err := r.GetUserItem(comment.UserId)

	item := &commentpb.CommentItem{
		Comment:  comment,
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

	return commentID, nil
}
