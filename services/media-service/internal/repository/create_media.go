package repository

import (
	"github.com/google/uuid"
	"github.com/salazarhugo/cheers1/services/media-service/internal/models"
)

func (m mediaRepository) CreateMedia(
	media *models.Media,
) (string, error) {
	db := m.spanner

	media.MediaId = uuid.NewString()

	result := db.Table("media").Create(&media)
	if result.Error != nil {
		return "", result.Error
	}

	return media.MediaId, nil
}
