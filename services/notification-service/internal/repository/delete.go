package repository

import "context"

func (r repository) DeleteToken(
	userID string,
	token string,
) error {
	ctx := context.Background()
	err := r.redis.SRem(ctx, userID, token, 0).Err()
	return err
}
