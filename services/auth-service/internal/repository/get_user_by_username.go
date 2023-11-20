package repository

import (
	"github.com/salazarhugo/cheers1/services/auth-service/internal/model"
)

func (a authRepository) GetUserByUsername(
	username string,
) (*model.User, error) {
	db := a.spanner
	var user model.User

	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
