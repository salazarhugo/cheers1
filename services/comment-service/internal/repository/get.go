package repository

import (
	"context"
	"encoding/json"
	"github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
	"log"
	"strings"
)

func (r repository) GetComment(commentId string) (*comment.CommentItem, error) {
	values, err := r.redis.ZScan(context.Background(), getKeyPostComment(postId), 0, 20).Result()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	com := &comment.Comment{}
	dec := json.NewDecoder(strings.NewReader(values[i]))
	err := dec.Decode(com)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	userItem, err := r.GetUserItem(com.UserId)

	item := &comment.CommentItem{
		Comment:  com,
		UserItem: userItem,
	}

	return item, nil
}
