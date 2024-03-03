package repository

import "github.com/salazarhugo/cheers1/libs/utils/models"

func (c chatRepository) GetUserNode(
	userID string,
) (*models.User, error) {
	db := c.spanner

	var user models.User

	result := db.
		Where("UserId = ?", userID).
		First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
