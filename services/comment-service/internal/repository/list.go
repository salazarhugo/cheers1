package repository

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
)

func (r repository) ListComment(
	postId string,
) ([]*comment.CommentItem, error) {
	values, err := r.redis.ZRevRange(context.Background(), getKeyPostComment(postId), 0, 20).Result()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var comments []*comment.CommentItem

	for i := range values {
		com := &comment.Comment{}
		err = protojson.Unmarshal([]byte(values[i]), com)
		userItem, err := r.GetUserItem(com.UserId)
		if err != nil {
			log.Println(err)
			continue
		}

		comments = append(comments, &comment.CommentItem{
			Comment:  com,
			UserItem: userItem,
		})
	}

	return comments, nil
}
