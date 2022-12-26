package cache

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
)

func (cache *redisCache) ListUser(
	userIds []string,
) ([]*user.UserItem, error) {
	items := make([]*user.UserItem, 0)

	for _, userId := range userIds {
		item, _ := cache.GetUserItem(userId)
		items = append(items, item)
	}

	return items, nil
}
