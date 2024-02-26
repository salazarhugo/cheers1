package tests

import (
	"github.com/salazarhugo/cheers1/libs/utils/models"
	"gorm.io/gorm"
)

func GetUserById(
	db *gorm.DB,
	userID string,
) (*models.User, error) {
	var user models.User

	result := db.
		Where("UserId = ?", userID).
		Or("username = ?", userID).
		First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
