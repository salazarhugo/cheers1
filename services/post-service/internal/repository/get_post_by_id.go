package repository

func (p *postRepository) GetPostById(
	postID string,
) (*Post, error) {
	db := p.spanner
	var post Post

	result := db.
		Table("posts").
		Where("PostId = ?", postID).
		First(&post)

	if result.Error != nil {
		return &Post{}, result.Error
	}

	return &post, nil
}
