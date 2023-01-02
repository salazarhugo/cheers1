package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	commentpb "github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"log"
	"time"
)

func getKeyPostComment(postId string) string {
	return fmt.Sprintf("%s:%s", keyPostComment, postId)
}

func (r repository) CreateComment(
	userId string,
	text string,
	postId string,
) error {
	id := uuid.New().String()
	comment := map[string]interface{}{
		"id":          id,
		"text":        text,
		"user_id":     userId,
		"create_time": time.Now().Unix(),
		"post_id":     postId,
	}

	buff := bytes.NewBufferString("")
	enc := json.NewEncoder(buff)
	err := enc.Encode(comment)
	if err != nil {
		return err
	}

	err = r.redis.ZAdd(
		context.Background(),
		getKeyPostComment(postId),
		redis.Z{
			Score:  float64(time.Now().Unix()),
			Member: buff.String(),
		},
	).Err()

	if err != nil {
		return err
	}
	commentItem, err := r.GetComment(id)

	go func() {
		err := pubsub.PublishProtoWithBinaryEncoding("comment-topic", &commentpb.CommentEvent{
			Event: &commentpb.CommentEvent_Created{
				Created: &commentpb.CreatedComment{
					Comment: commentItem.Comment,
					User:    commentItem.UserItem,
				},
			},
		})
		if err != nil {
			log.Println(err)
		}
	}()

	return nil
}
