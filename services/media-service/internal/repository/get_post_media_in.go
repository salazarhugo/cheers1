package repository

import "github.com/salazarhugo/cheers1/services/media-service/internal/models"

func (m mediaRepository) ListPostMedia(
	postID string,
) ([]*models.PostMedia, error) {
	db := m.spanner

	var postMedias []*models.PostMedia

	result := db.
		Table("post_media").
		Where("PostId = ?", postID).
		Scan(&postMedias)

	if result.Error != nil {
		return nil, result.Error
	}

	return postMedias, nil
}
