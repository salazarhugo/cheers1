package repository

import (
	"github.com/salazarhugo/cheers1/libs/utils/models"
)

func (p *userRepository) UpdateUser(
	user *models.User,
) (string, error) {
	db := p.spanner

	result := db.Updates(&user)
	if result.Error != nil {
		return "", result.Error
	}

	return user.UserId, nil
}
