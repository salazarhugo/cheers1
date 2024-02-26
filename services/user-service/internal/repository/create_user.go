package repository

import (
	"github.com/salazarhugo/cheers1/libs/utils/models"
)

func (p *userRepository) CreateUser(
	userID string,
	username string,
	name string,
	picture string,
	email string,
) (string, error) {
	db := p.spanner

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
