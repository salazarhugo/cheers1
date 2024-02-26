package tests

import (
	"github.com/salazarhugo/cheers1/libs/utils/tests"
	"github.com/salazarhugo/cheers1/services/note-service/internal/repository"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFeedNote(t *testing.T) {
	//Create a mock repository
	repo := repository.NewRepository()

	user := tests.CreateRandomUser(t)

	_, err := repo.FeedNote(
		user.UserId,
		1,
		10,
	)
	require.NoError(t, err)
}
