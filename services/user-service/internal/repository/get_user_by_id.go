package repository

import (
	"github.com/salazarhugo/cheers1/services/user-service/internal/model"
)

func (p *userRepository) GetUserById(
	userID string,
) (model.User, error) {
	db := p.spanner
	var user model.User

	result := db.Where("id = ?", userID).Or("username = ?", userID).First(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}

	return user, nil
}
