package repository

import "github.com/salazarhugo/cheers1/libs/utils/models"

func (p *postRepository) DeletePost(
	postID string,
) error {
	db := p.spanner
	result := db.Delete(&models.Post{PostId: postID})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
