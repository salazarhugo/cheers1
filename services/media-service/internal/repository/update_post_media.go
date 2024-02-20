package repository

import (
	"github.com/salazarhugo/cheers1/services/media-service/internal/models"
)

func (m mediaRepository) UpdatePostMedia(
	postMedia *models.PostMedia,
) error {
	db := m.spanner

	result := db.
		Table("post_media").
		Updates(&postMedia)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
