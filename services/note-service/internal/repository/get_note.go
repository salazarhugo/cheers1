package repository

import (
	"github.com/salazarhugo/cheers1/libs/utils/models"
)

func (r *repository) GetNote(
	userID string,
) (*models.UserStatus, error) {
	db := r.spanner
	var userStatus models.UserStatus

	result := db.
		Table("user_status").
		Where("UserId = ?", userID).
		First(&userStatus)

	if result.Error != nil {
		return nil, result.Error
	}

	return &userStatus, nil
}
