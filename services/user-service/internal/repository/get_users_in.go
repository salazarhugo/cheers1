package repository

import (
	"github.com/salazarhugo/cheers1/libs/utils/models"
)

func (p *userRepository) GetUsersIn(
	userIDs []string,
) ([]*models.User, error) {
	db := p.spanner

	var users []*models.User

	result := db.
		Where("username IN ?", userIDs).
		Or("UserId IN ?", userIDs).
		Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
