package tests

import (
	"github.com/salazarhugo/cheers1/services/note-service/internal/repository"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeleteNote(t *testing.T) {
	repo := repository.NewRepository()

	note := CreateRandomNote(t)

	err := repo.DeleteNote(note.UserId)
	require.NoError(t, err)

	notes, err := repo.FeedNote(
		note.UserId,
		1,
		10,
	)
	require.NoError(t, err)
	require.Empty(t, notes)
}
