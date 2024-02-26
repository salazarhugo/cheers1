package repository

import "github.com/salazarhugo/cheers1/libs/utils/models"

func (p *postRepository) GetPostById(
	postID string,
) (*models.Post, error) {
	db := p.spanner
	var post models.Post

	result := db.
		Table("posts").
		Where("PostId = ?", postID).
		First(&post)

	if result.Error != nil {
		return &models.Post{}, result.Error
	}

	return &post, nil
}
