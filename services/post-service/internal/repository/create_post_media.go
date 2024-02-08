package repository

import (
	"github.com/google/uuid"
	"github.com/salazarhugo/cheers1/services/post-service/internal/models"
	"gorm.io/gorm"
)

func (p *postRepository) CreatePostMedia(
	tx *gorm.DB,
	postMedia *models.PostMedia,
) *gorm.DB {
	postMedia.PostMediaId = uuid.NewString()

	result := tx.
		Table("post_media").
		Create(&postMedia)

	return result
}
