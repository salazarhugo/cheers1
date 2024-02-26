package repository

import (
	"github.com/salazarhugo/cheers1/libs/utils/models"
)

func (p *userRepository) GetUserById(
	userID string,
) (models.User, error) {
	db := p.spanner
	var user models.User

	result := db.
		Where("UserId = ?", userID).
		Or("username = ?", userID).
		First(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}
