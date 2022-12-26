package cache

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
)

func (cache *redisCache) UpdateUser(user *user.User) error {
	err := cache.client.HSet(context.Background(), getKeyUser(user.Id), map[string]interface{}{
		"name":     user.Name,
		"username": user.Username,
		"picture":  user.Picture,
		"verified": user.Verified,
	}).Err()

	return err
}
