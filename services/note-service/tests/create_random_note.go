package tests

import (
	note2 "github.com/salazarhugo/cheers1/gen/go/cheers/note/v1"
	"github.com/salazarhugo/cheers1/libs/utils/models"
	"github.com/salazarhugo/cheers1/libs/utils/tests"
	"github.com/salazarhugo/cheers1/services/note-service/internal/domain"
	"github.com/salazarhugo/cheers1/services/note-service/internal/repository"
	"github.com/stretchr/testify/require"
	"testing"
)

func CreateRandomNote(t *testing.T) *models.UserStatus {
	usecases := domain.NewDomain()
	repo := repository.NewRepository()

	user := tests.CreateRandomUser(t)

	params := domain.CreateNoteUseCaseParams{
		UserID:   user.UserId,
		NoteType: note2.NoteType_NOTHING,
		Text:     "drinking at my favorite bar",
		DrinkID:  nil,
	}
	_, err := usecases.CreateNoteUseCase(params)
	require.NoError(t, err)

	note, err := repo.GetNote(user.UserId)
	require.NoError(t, err)
	require.Equal(t, user.UserId, note.UserId)
	require.Equal(t, params.NoteType.String(), note.Type)
	require.Equal(t, params.Text, note.Text)

	return note
}
