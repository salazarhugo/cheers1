package tests

import (
	"github.com/google/uuid"
	"github.com/salazarhugo/cheers1/libs/utils/tests"
	"github.com/salazarhugo/cheers1/services/user-service/internal/model"
	"github.com/salazarhugo/cheers1/services/user-service/internal/repository"
	"github.com/stretchr/testify/require"
	"testing"
)

func CreateRandomUser(t *testing.T) *model.User {
	//Create a mock repository
	repo := repository.NewUserRepository()

	// GIVEN
	userName := "User Service"

	userID, err := repo.CreateUser(
		uuid.NewString(),
		tests.CreateRandomUsername(10),
		userName,
		"",
		"test2@cheers.social",
	)

	require.NoError(t, err)

	user, err := repo.GetUserById(userID)

	require.NoError(t, err)
	require.Equal(t, userName, user.Name)

	// CLEANUP
	t.Cleanup(func() {
		err = repo.DeleteUser(user.UserId)
		require.NoError(t, err)
	})

	return &user
}
