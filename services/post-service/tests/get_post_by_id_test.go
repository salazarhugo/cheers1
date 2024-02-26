package tests

import (
	"github.com/salazarhugo/cheers1/services/post-service/internal/repository"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetPostById(t *testing.T) {
	// GIVEN
	repo := repository.NewRepository()
	post := CreateRandomPost(t)

	// TEST
	response, err := repo.GetPostById(
		post.PostId,
	)
	require.NoError(t, err)
	require.NotNil(t, response)
}
