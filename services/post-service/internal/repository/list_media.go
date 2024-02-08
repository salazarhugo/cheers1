package repository

import "github.com/salazarhugo/cheers1/services/post-service/internal/models"

func (p *postRepository) GetMedias(mediaIDs []string) ([]*models.Media, error) {
	var medias []*models.Media

	result := p.spanner.
		Table("media").
		Where("media.MediaId IN ?", mediaIDs).
		Scan(&medias)

	if result.Error != nil {
		return nil, result.Error
	}

	return medias, nil
}
