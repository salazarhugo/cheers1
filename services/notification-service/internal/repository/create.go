package repository

import "context"

func (r repository) CreateRegistrationToken(
	userID string,
	token string,
) error {
	ctx := context.Background()
	err := r.redis.SAdd(ctx, userID, token, 0).Err()
	return err
}
