package tests

import (
	"github.com/google/uuid"
	"github.com/salazarhugo/cheers1/services/user-service/internal/repository"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCheckUsernameNotExist(t *testing.T) {
	//Create a mock repository
	repo := repository.NewUserRepository()

	username := uuid.NewString()
	exists, err := repo.CheckUsername(username)

	require.NoError(t, err)
	require.False(t, exists)
}

func TestCheckUsernameExist(t *testing.T) {
	user := CreateRandomUser(t)

	//Create a mock repository
	repo := repository.NewUserRepository()

	username := user.Username
	exists, err := repo.CheckUsername(username)

	require.NoError(t, err)
	require.True(t, exists)
}
