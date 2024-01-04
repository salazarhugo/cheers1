package repository

func (p *postRepository) GetPostTotalLikes(
	postID string,
) (int64, error) {
	db := p.spanner
	var total int64

	result := db.Table("post_likes").
		Where("post_id = ?", postID).
		Count(&total)
	if result.Error != nil {
		return 0, result.Error
	}

	return total, nil
}
