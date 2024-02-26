package repository

import (
	"github.com/salazarhugo/cheers1/services/user-service/internal/model"
)

func (p *userRepository) DeleteUser(
	userID string,
) error {
	db := p.spanner
	result := db.Delete(&model.User{UserId: userID})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
