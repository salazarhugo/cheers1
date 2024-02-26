package repository

import (
	"github.com/salazarhugo/cheers1/services/user-service/internal/model"
)

func (p *userRepository) UpdateUser(
	user *model.User,
) (string, error) {
	db := p.spanner

	result := db.Updates(&user)
	if result.Error != nil {
		return "", result.Error
	}

	return user.UserId, nil
}
