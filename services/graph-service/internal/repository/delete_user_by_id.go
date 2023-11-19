package repository

import (
	"github.com/salazarhugo/cheers1/services/user-service/internal/model"
)

func (p *userRepository) DeleteUserById(
	userID string,
) error {
	db := p.spanner
	result := db.Delete(&model.User{ID: userID})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
