package tests

import (
	"github.com/salazarhugo/cheers1/libs/utils/tests"
	"github.com/salazarhugo/cheers1/services/friendship-service/internal/repository"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestListFriend(t *testing.T) {
	//Create a mock repository
	repo := repository.NewRepository()

	// Given
	user1 := tests.CreateRandomUser(t)
	user2 := tests.CreateRandomUser(t)
	err := repo.CreateFriend(user1.UserId, user2.UserId)
	require.NoError(t, err)

	// Test
	friendsOfUser1, err := repo.ListFriend(
		user1.UserId,
		1,
		100,
		user1.UserId,
	)
	require.NoError(t, err)
	require.Len(t, friendsOfUser1, 1)
	require.Equal(t, user2.Username, friendsOfUser1[0].Username)
	require.True(t, friendsOfUser1[0].Friend)

	friendsOfUser2, err := repo.ListFriend(
		user1.UserId,
		1,
		100,
		user2.UserId,
	)
	require.NoError(t, err)
	require.Len(t, friendsOfUser2, 1)
	require.Equal(t, user1.Username, friendsOfUser2[0].Username)
	require.False(t, friendsOfUser2[0].Friend)
}
