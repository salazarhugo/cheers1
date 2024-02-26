package tests

import (
	"github.com/salazarhugo/cheers1/services/user-service/internal/repository"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSearchUser(t *testing.T) {
	//Create a mock repository
	repo := repository.NewUserRepository()
	user1 := CreateRandomUser(t)

	// Remove last character of username
	query := user1.Username[:len(user1.Username)-1]

	users, err := repo.SearchUser(query)
	require.NoError(t, err)

	require.GreaterOrEqual(t, len(users), 1)
	require.Contains(t, users, user1)
}
