package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/note/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
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

	ListNote(
		userID string,
	) ([]*note.Note, error)
}

type repository struct {
	firestore *firestore.Client
}

func NewRepository() Repository {
	fire, err := utils.InitializeAppDefault().Firestore(context.Background())
	if err != nil {
		return nil
	}

	return &repository{
		firestore: fire,
	}
}
