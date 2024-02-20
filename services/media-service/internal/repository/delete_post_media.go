package repository

import "github.com/salazarhugo/cheers1/services/media-service/internal/models"

func (m mediaRepository) DeletePostMedia(
	postMediaID string,
) error {
	db := m.spanner

	result := db.
		Table("post_media").
		Delete(&models.PostMedia{PostMediaId: postMediaID})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
