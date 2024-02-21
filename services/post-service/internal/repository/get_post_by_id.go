package repository

func (p *postRepository) GetPostById(
	postID string,
) (*Post, error) {
	db := p.spanner
	var post Post

	mediaQuery := db.
		Raw("SELECT TO_JSON(t) FROM post_media AS t WHERE posts.PostId = t.PostId")

	result := db.
		Table("posts").
		Select(
			"posts.*, ARRAY(?) AS medias",
			mediaQuery,
		).
		Where("PostId = ?", postID).
		First(&post)
	if result.Error != nil {
		return &Post{}, result.Error
	}

	return &post, nil
}
