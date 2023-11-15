package repository

func (p *postRepository) DeletePostById(
	postID string,
) error {
	db := p.spanner
	result := db.Delete(&Post{ID: postID})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
