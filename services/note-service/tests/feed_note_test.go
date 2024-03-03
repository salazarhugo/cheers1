package tests

import (
	"github.com/salazarhugo/cheers1/services/note-service/internal/repository"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFeedNote(t *testing.T) {
	//Create a mock repository
	repo := repository.NewRepository()

	note := CreateRandomNote(t)

	notes, err := repo.FeedNote(
		note.UserId,
		1,
		10,
	)
	require.NoError(t, err)
	require.Len(t, notes, 1)
	t.Log(notes[0])
}
