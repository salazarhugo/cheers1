package repository

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"github.com/salazarhugo/cheers1/services/notification-service/internal/notifications"
)

func (r repository) FollowUserNotification(
	user *user.User,
	followedUser *user.User,
) error {

	data := notifications.FollowUserNotification(user.Username, user.Picture)

	tokens, err := r.GetUserTokens(followedUser.Id)
	if err != nil {
		return err
	}

	err = r.SendNotification(
		map[string][]string{followedUser.Id: tokens},
		data,
	)

	if err != nil {
		return err
	}

	return nil
}