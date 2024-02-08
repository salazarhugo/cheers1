package repository

import (
	"github.com/salazarhugo/cheers1/services/post-service/internal/models"
	"gorm.io/gorm"
)

func (p *postRepository) DeleteMedia(
	tx *gorm.DB,
	mediaID string,
) *gorm.DB {
	result := tx.
		Table("media").
		Delete(&models.Media{MediaId: mediaID})

	return result
}
