package repository

import (
	"github.com/salazarhugo/cheers1/services/user-service/internal/model"
)

func (p *userRepository) GetUserByUsername(
	username string,
) (model.User, error) {
	db := p.spanner
	var user model.User

	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}

	return user, nil
}
