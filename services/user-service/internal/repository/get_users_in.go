package repository

import (
	"github.com/salazarhugo/cheers1/services/user-service/internal/model"
)

func (p *userRepository) GetUsersIn(
	userIDs []string,
) ([]*model.User, error) {
	db := p.spanner

	var users []*model.User

	result := db.
		Where("username IN ?", userIDs).
		Or("UserId IN ?", userIDs).
		Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
