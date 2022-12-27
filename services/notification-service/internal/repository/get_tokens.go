package repository

import (
	"context"
)

func (r repository) GetUserTokens(userID string) ([]string, error) {
	ctx := context.Background()
	tokens, err := r.redis.SMembers(ctx, userID).Result()

	if err != nil {
		return nil, err
	}

	return tokens, nil
}
