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

func (r repository) GetUsersTokens(userIDs []string) (map[string][]string, error) {
	ctx := context.Background()

	tokens := make(map[string][]string, 0)

	for _, userID := range userIDs {
		userTokens, err := r.redis.SMembers(ctx, userID).Result()
		if err != nil {
			return nil, err
		}
		tokens[userID] = userTokens
	}

	return tokens, nil
}
