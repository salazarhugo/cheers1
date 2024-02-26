package repository

import (
	"fmt"
	"github.com/salazarhugo/cheers1/libs/utils/models"
)

func (p *userRepository) SearchUser(
	query string,
) ([]*models.User, error) {
	db := p.spanner

	var users []*models.User
	// Adds prefix and suffix '%' to query => %cheers%
	query = fmt.Sprintf("%%%s%%", query)

	result := db.
		Table("users").
		Where("username LIKE ?", query).
		Or("name LIKE ?", query).
		Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
