package tests

import (
	"github.com/salazarhugo/cheers1/libs/utils/tests"
	"github.com/salazarhugo/cheers1/services/friendship-service/internal/repository"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateFriendRequest(t *testing.T) {
	//Create a mock repository
	repo := repository.NewRepository()

	// Given
	user1 := tests.CreateRandomUser(t)
	user2 := tests.CreateRandomUser(t)

	// Test
	err := repo.CreateFriendRequest(user1.UserId, user2.UserId)
	require.NoError(t, err)

	friendRequests, err := repo.ListFriendRequests(1, 10, user2.UserId)
	require.NoError(t, err)
	require.Len(t, friendRequests, 1)
}
