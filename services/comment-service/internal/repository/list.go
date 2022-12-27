package repository

import (
	"context"
	"encoding/json"
	"github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
	"log"
	"strings"
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
		dec := json.NewDecoder(strings.NewReader(values[i]))
		err := dec.Decode(com)
		if err != nil {
			log.Println(err)
			continue
		}

		userItem, err := r.GetUserItem(com.UserId)

		comments = append(comments, &comment.CommentItem{
			Comment:  com,
			UserItem: userItem,
		})
	}

	return comments, nil
}
