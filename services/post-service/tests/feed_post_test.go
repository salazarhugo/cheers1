package tests

import (
	"github.com/salazarhugo/cheers1/services/post-service/internal/repository"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFeedPost(t *testing.T) {
	//Create a mock repository
	repo := repository.NewRepository()

	post := CreateRandomPost(t)

	response, err := repo.FeedPost(
		[]string{post.UserID},
		1,
		10,
	)
	require.NoError(t, err)

	require.Len(t, response.GetPosts(), 1)
}
