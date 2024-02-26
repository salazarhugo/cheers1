package tests

import (
	"github.com/google/uuid"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func CreateRandomUser(t *testing.T) *models.User {
	// GIVEN
	db := utils.GetSpanner()
	userName := "User Service"

	userID, err := CreateUser(
		db,
		uuid.NewString(),
		CreateRandomUsername(10),
		userName,
		"",
		"test2@cheers.social",
	)

	require.NoError(t, err)

	user, err := GetUserById(db, userID)

	require.NoError(t, err)
	require.Equal(t, userName, user.Name)

	// CLEANUP
	t.Cleanup(func() {
		err = DeleteUser(db, user.UserId)
		require.NoError(t, err)
	})

	return user
}
