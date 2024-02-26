package repository

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/note/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"gorm.io/gorm"
)

type Repository interface {
	CreateNote(
		userID string,
		text string,
	) (string, error)

	GetNote(
		userID string,
	) (*note.Note, error)

	DeleteNote(
		userID string,
	) error

	FeedNote(
		userID string,
		page int,
		pageSize int,
	) ([]*note.Note, error)
}

type repository struct {
	spanner *gorm.DB
}

func NewRepository() Repository {
	return &repository{
		spanner: utils.GetSpanner(),
	}
}
