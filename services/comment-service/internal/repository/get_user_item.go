package repository

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
)

func (r repository) GetUserItem(
	userId string,
) (*user.UserItem, error) {
	res, err := r.redis.HGetAll(context.Background(), getKeyUser(userId)).Result()
	if err != nil {
		return nil, err
	}
	item := &user.UserItem{
		Id:          res["id"],
		Name:        res["name"],
		Username:    res["username"],
		Verified:    res["verified"] == "1",
		Picture:     res["picture"],
		HasFollowed: false,
		StoryState:  0,
	}

	return item, nil
}
