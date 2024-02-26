package tests

import (
	"github.com/salazarhugo/cheers1/services/post-service/internal/repository"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetPostItem(t *testing.T) {
	//Create a mock repository
	repo := repository.NewRepository()

	post := CreateRandomPost(t)

	response, err := repo.GetPostItem(
		post.UserID,
		post.PostId,
	)
	require.NoError(t, err)
	require.Zero(t, response.LikeCount)
	require.Zero(t, response.CommentCount)
	require.False(t, response.HasLiked)
}
