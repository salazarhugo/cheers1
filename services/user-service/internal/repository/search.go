package repository

import (
	"fmt"
	"github.com/salazarhugo/cheers1/services/user-service/internal/model"
)

func (p *userRepository) SearchUser(
	query string,
) ([]*model.User, error) {
	db := p.spanner

	var users []*model.User
	// Adds prefix and suffix '%' to query => %cheers%
	query = fmt.Sprintf("%%%s%%", query)

	result := db.Table("users").Where("username LIKE ?", query).Or("name LIKE ?", query).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
