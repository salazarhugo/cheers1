package repository

import (
	"github.com/salazarhugo/cheers1/libs/utils/models"
)

func (p *userRepository) DeleteUser(
	userID string,
) error {
	db := p.spanner
	result := db.Delete(&models.User{UserId: userID})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
