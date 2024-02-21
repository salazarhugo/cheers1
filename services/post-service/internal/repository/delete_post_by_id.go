package repository

func (p *postRepository) DeletePost(
	postID string,
) error {
	db := p.spanner
	result := db.Delete(&Post{PostId: postID})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
