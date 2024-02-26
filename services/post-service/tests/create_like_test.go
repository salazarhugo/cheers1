package tests

import (
	"github.com/salazarhugo/cheers1/services/post-service/internal/repository"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateLike(t *testing.T) {
	//Create a mock repository
	repo := repository.NewRepository()

	post := CreateRandomPost(t)

	_, err := repo.LikePost(
		post.UserID,
		post.PostId,
	)
	require.NoError(t, err)

	_, err = repo.UnlikePost(
		post.UserID,
		post.PostId,
	)
	require.NoError(t, err)
}
