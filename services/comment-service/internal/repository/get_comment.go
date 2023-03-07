package repository

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
)

func (r repository) GetComment(
	commentID string,
) (*comment.Comment, error) {
	ctx := context.Background()

	com := &comment.Comment{}
	commentFields, err := r.redis.HGetAll(ctx, getKeyComment(commentID)).Result()
	if err != nil {
		return nil, err
	}
	err = utils.MapToProto(com, commentFields)
	if err != nil {
		return nil, err
	}

	return com, nil
}
