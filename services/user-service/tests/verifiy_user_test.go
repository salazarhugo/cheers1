package tests

import (
	"github.com/salazarhugo/cheers1/services/user-service/internal/repository"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVerifyUser(t *testing.T) {
	//Create a mock repository
	repo := repository.NewUserRepository()

	user := CreateRandomUser(t)

	err := repo.VerifyUser(
		user.UserId,
		true,
	)
	require.NoError(t, err)

	updatedUser, err := repo.GetUserById(user.UserId)
	require.NoError(t, err)
	require.NotNil(t, updatedUser)

	require.True(t, updatedUser.Verified)
}
