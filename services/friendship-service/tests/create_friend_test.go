package tests

import (
	"github.com/salazarhugo/cheers1/libs/utils/tests"
	"github.com/salazarhugo/cheers1/services/friendship-service/internal/repository"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateFriend(t *testing.T) {
	//Create a mock repository
	repo := repository.NewRepository()

	// Given
	user1 := tests.CreateRandomUser(t)
	user2 := tests.CreateRandomUser(t)

	// Test
	err := repo.CreateFriend(user1.UserId, user2.UserId)
	require.NoError(t, err)
}
