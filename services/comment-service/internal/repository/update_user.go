package repository

import (
	"context"
	"fmt"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
)

const (
	keyUser     = "user"
	keyPost     = "post"
	keyComment  = "comment"
	keyComments = "comments"
	keyReplies  = "replies"
	keyLikes    = "likes"
)

func getKeyUser(userUUID string) string {
	return fmt.Sprintf("%s:%s", keyUser, userUUID)
}

func (r repository) UpdateUser(user *user.User) error {
	err := r.redis.HSet(context.Background(), getKeyUser(user.Id), map[string]interface{}{
		"id":       user.Id,
		"name":     user.Name,
		"username": user.Username,
		"picture":  user.Picture,
		"verified": user.Verified,
	}).Err()

	return err
}
