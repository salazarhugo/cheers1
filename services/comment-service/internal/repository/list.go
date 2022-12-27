package repository

import (
	"context"
	"encoding/json"
	"github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
	"strings"
)

func (r repository) ListComment(postId string) ([]*comment.Comment, error) {
	values, err := r.redis.ZRevRange(context.Background(), getKeyPostComment(postId), 0, 20).Result()
	if err != nil {
		return nil, err
	}

	var comments []*comment.Comment

	for i := range values {
		comment := &comment.Comment{}
		dec := json.NewDecoder(strings.NewReader(values[i]))
		err := dec.Decode(comment)
		if err != nil {
			panic(err)
		}
		comments = append(comments, comment)
	}

	return comments, nil
}
