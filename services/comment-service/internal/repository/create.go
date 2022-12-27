package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"time"
)

func getKeyPostComment(postId string) string {
	return fmt.Sprintf("%s:%s", keyPostComment, postId)
}

func (r repository) CreateComment(
	text string,
	postId string,
) error {
	comment := map[string]interface{}{
		"id":         uuid.New().String(),
		"text":       text,
		"createTime": time.Now().Unix(),
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

	return nil
}
