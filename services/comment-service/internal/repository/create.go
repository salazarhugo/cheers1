package repository

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	commentpb "github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
	"github.com/salazarhugo/cheers1/libs/utils/pubsub"
	"google.golang.org/protobuf/encoding/protojson"
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
	comment := &commentpb.Comment{
		Id:         uuid.New().String(),
		Text:       text,
		CreateTime: time.Now().Unix(),
		UserId:     userId,
		PostId:     postId,
	}

	bytes, err := protojson.Marshal(comment)

	err = r.redis.ZAdd(
		context.Background(),
		getKeyPostComment(postId),
		redis.Z{
			Score:  float64(time.Now().Unix()),
			Member: string(bytes),
		},
	).Err()

	if err != nil {
		return err
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
					Comment: item.Comment,
					User:    item.UserItem,
				},
			},
		})
		if err != nil {
			log.Println(err)
		}
	}()

	return nil
}
