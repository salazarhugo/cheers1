package tests

import (
	"github.com/salazarhugo/cheers1/libs/utils/models"
	"gorm.io/gorm"
)

func CreateUser(
	db *gorm.DB,
	userID string,
	username string,
	name string,
	picture string,
	email string,
) (string, error) {
	user := &models.User{}
	user.UserId = userID
	user.Username = username
	user.Name = name
	user.Picture = picture
	user.Email = email
	user.Verified = false

	result := db.Create(&user)
	if result.Error != nil {
		return "", result.Error
	}

	return userID, nil
}
