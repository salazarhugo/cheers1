package repository

import (
	"github.com/salazarhugo/cheers1/services/user-service/internal/model"
)

func (p *userRepository) CreateUser(
	userID string,
	username string,
	name string,
	picture string,
	email string,
) (string, error) {
	db := p.spanner

	user := &model.User{}
	user.ID = userID
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
