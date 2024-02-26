package tests

import (
	"github.com/salazarhugo/cheers1/libs/utils/models"
	"gorm.io/gorm"
)

func DeleteUser(
	db *gorm.DB,
	userID string,
) error {
	result := db.Delete(&models.User{UserId: userID})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
