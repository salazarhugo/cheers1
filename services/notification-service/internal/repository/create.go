package repository

import "context"

func (r repository) CreateRegistrationToken(
	userID string,
	token string,
) error {
	ctx := context.Background()
	err := r.redis.SAdd(ctx, userID, token).Err()
	return err
}
