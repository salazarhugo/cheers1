package repository

import (
	uuid "github.com/google/uuid"
)

func (p *postRepository) CreatePost(
	userID string,
	post *Post,
) (string, error) {
	db := p.spanner

	post.PostId = uuid.NewString()
	post.UserID = userID

	result := db.Table("posts").Create(&post)
	if result.Error != nil {
		return "", result.Error
	}

	return post.PostId, nil
}
