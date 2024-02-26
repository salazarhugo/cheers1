package repository

import (
	uuid "github.com/google/uuid"
	"github.com/salazarhugo/cheers1/libs/utils/models"
)

func (p *postRepository) CreatePost(
	userID string,
	post *models.Post,
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
