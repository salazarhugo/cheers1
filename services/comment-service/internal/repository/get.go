package repository

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
)

func (repository repository) GetLastComment(postId string) (*comment.CommentItem, error) {
	values, err := repository.redis.ZRevRangeByScore(context.Background(), getKeyPostComment(postId), &redis.ZRangeBy{
		Offset: 0,
		Count:  1,
	}).Result()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if len(values) < 1 {
		return nil, nil
	}

	com := &comment.Comment{}
	err = protojson.Unmarshal([]byte(values[0]), com)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	userItem, err := repository.GetUserItem(com.UserId)

	item := &comment.CommentItem{
		Comment:  com,
		UserItem: userItem,
	}

	return item, nil
}
