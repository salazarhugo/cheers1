package tests

import (
	"github.com/salazarhugo/cheers1/services/user-service/internal/model"
	"github.com/salazarhugo/cheers1/services/user-service/internal/repository"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetUsersIn(t *testing.T) {
	//Create a mock repository
	repo := repository.NewUserRepository()

	user1 := CreateRandomUser(t)
	user2 := CreateRandomUser(t)

	users := []*model.User{user1, user2}

	query := []string{user1.Username, user2.Username}

	res, err := repo.GetUsersIn(query)
	require.NoError(t, err)
	require.Equal(t, len(users), len(res))
}
