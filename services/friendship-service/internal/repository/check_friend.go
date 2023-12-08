package repository

import (
	"context"
)

func (r repository) CheckFriend(user1 string, user2 string) (bool, error) {
	ctx := context.Background()
	isUser1FriendWithUser2, err := r.redis.SIsMember(ctx, getKeyFriends(user1), user2).Result()
	if err != nil {
		return false, err
	}

	if isUser1FriendWithUser2 == false {
		return false, nil
	}

	isUser2FriendWithUser1, err := r.redis.SIsMember(ctx, getKeyFriends(user2), user1).Result()
	if err != nil {
		return false, err
	}

	if isUser2FriendWithUser1 == false {
		return false, nil
	}

	return true, err
}
