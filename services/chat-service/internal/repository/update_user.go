package repository

import "github.com/salazarhugo/cheers1/gen/go/cheers/type/user"

func (c chatRepository) UpdateUser(user *user.User) error {
	if user == nil {
		return nil
	}
	return c.cache.UpdateUser(user)
}
