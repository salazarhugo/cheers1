package repository

import (
	"context"
)

func (r repository) UpdateGhostMode(
	userId string,
	ghostMode bool,
) error {
	ctx := context.Background()

	if ghostMode {
		return r.redis.SAdd(ctx, keyGhostMode, userId).Err()
	} else {
		return r.redis.SRem(ctx, keyGhostMode, userId).Err()
	}
}
