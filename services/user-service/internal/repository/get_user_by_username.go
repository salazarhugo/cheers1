package repository

import (
	"github.com/salazarhugo/cheers1/libs/utils/models"
)

func (p *userRepository) GetUserByUsername(
	username string,
) (*models.User, error) {
	db := p.spanner
	var user models.User

	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
