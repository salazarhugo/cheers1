package repository

import (
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/media-service/internal/models"
	"gorm.io/gorm"
)

type MediaRepository interface {
	CreateMedia(media *models.Media) (string, error)
	UpdatePostMedia(postMedia *models.PostMedia) error
	DeletePostMedia(postMediaID string) error
	ListPostMedia(postID string) ([]*models.PostMedia, error)
}

type mediaRepository struct {
	spanner *gorm.DB
}

func NewPostRepository() MediaRepository {
	return &mediaRepository{
		spanner: utils.GetSpanner(),
	}
}
