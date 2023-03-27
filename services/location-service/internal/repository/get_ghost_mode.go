package repository

import (
	"context"
)

func (r repository) GetGhostMode(
	userId string,
) (bool, error) {
	ctx := context.Background()

	res, err := r.redis.SIsMember(ctx, keyGhostMode, userId).Result()
	if err != nil {
		return false, err
	}

	return res, err
}
