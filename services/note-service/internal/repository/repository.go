package repository

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/note/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateNote(
		status *models.UserStatus,
	) (string, error)

	GetNote(
		userID string,
	) (*models.UserStatus, error)

	DeleteNote(
		userID string,
	) error

	FeedNote(
		userID string,
		page int,
		pageSize int,
	) ([]*note.Note, error)

	ListFriendIds(
		userID string,
	) ([]string, error)
}

type repository struct {
	spanner *gorm.DB
}

func NewRepository() Repository {
	return &repository{
		spanner: utils.GetSpanner(),
	}
}
