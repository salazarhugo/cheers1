package repository

import "github.com/salazarhugo/cheers1/libs/utils/models"

func (c chatRepository) GetUsersNode(
	userIDs []string,
) ([]*models.User, error) {
	db := c.spanner

	var users []*models.User

	result := db.
		Table("users").
		Where("UserId IN ?", userIDs).
		Scan(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
