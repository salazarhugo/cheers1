package repository

import (
	"gorm.io/gorm"
)

func (p *postRepository) CreatePostMedias(
	postID string,
	mediaIDs []string,
) error {
	medias, err := p.GetMedias(mediaIDs)
	if err != nil {
		return err
	}

	for _, media := range medias {
		err := p.spanner.Transaction(func(tx *gorm.DB) error {
			if err := p.DeleteMedia(tx, media.MediaId).Error; err != nil {
				return err
			}

			if err := p.CreatePostMedia(tx, media.ToPostMedia(postID)).Error; err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
