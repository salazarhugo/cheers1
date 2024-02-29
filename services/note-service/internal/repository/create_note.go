package repository

import (
	"github.com/google/uuid"
	"github.com/salazarhugo/cheers1/libs/utils/models"
)

func (r *repository) CreateNote(
	note *models.UserStatus,
) (string, error) {
	db := r.spanner

	note.UserStatusId = uuid.NewString()

	result := db.
		Table("user_status").
		Create(&note)
	if result.Error != nil {
		return "", result.Error
	}

	return note.UserStatusId, nil
}
