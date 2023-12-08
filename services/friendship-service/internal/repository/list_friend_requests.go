package repository

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"github.com/salazarhugo/cheers1/services/friendship-service/internal/service"
	"log"
)

func (r repository) ListFriendRequests(
	userId string,
) ([]*user.UserItem, error) {
	ctx := context.Background()
	userIds, err := r.redis.SMembers(ctx, getKeyFriendRequests(userId)).Result()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	users, err := service.GetUsers(userIds)

	return users, nil
}
