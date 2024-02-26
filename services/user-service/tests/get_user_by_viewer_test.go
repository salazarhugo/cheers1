package tests

import (
	"github.com/salazarhugo/cheers1/services/user-service/internal/repository"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetUserByViewer(t *testing.T) {
	// SETUP
	user := CreateRandomUser(t)

	// GIVEN
	repo := repository.NewUserRepository()
	viewerID := user.UserId
	userID := user.UserId

	// TEST
	res, err := repo.GetUserWithViewer(viewerID, userID)
	require.NoError(t, err)
	require.NotNil(t, res)
}
