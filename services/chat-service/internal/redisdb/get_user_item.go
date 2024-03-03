package redisdb

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
)

func (cache *redisCache) GetUserItem(
	userId string,
) (*user.UserItem, error) {
	item := &user.UserItem{
		Id:          "",
		Name:        "",
		Username:    "",
		Verified:    true,
		Picture:     "",
		HasFollowed: false,
		StoryState:  0,
	}

	return item, nil
}
