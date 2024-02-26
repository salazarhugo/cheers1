package tests

import (
	"github.com/salazarhugo/cheers1/services/user-service/internal/repository"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeleteUser(t *testing.T) {
	repo := repository.NewUserRepository()

	user := CreateRandomUser(t)

	err := repo.DeleteUser(user.UserId)
	require.NoError(t, err)
}
