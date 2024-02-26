package tests

import (
	"cloud.google.com/go/spanner"
	"github.com/salazarhugo/cheers1/libs/utils/models"
	"github.com/salazarhugo/cheers1/libs/utils/tests"
	"github.com/salazarhugo/cheers1/services/post-service/internal/repository"
	"github.com/stretchr/testify/require"
	"testing"
)

func CreateRandomPost(t *testing.T) *models.Post {
	//Create a mock repository
	repo := repository.NewRepository()
	user := tests.CreateRandomUser(t)

	// GIVEN
	postRequest := &models.Post{
		Caption: "Only dead fish go with the flow",
		DrinkID: spanner.NullString{Valid: false},
	}
	postID, err := repo.CreatePost(
		user.UserId,
		postRequest,
	)
	require.NoError(t, err)

	post, err := repo.GetPostById(postID)
	require.NoError(t, err)
	require.Equal(t, postRequest.Caption, post.Caption)

	// CLEANUP
	t.Cleanup(func() {
		err = repo.DeletePost(postID)
		require.NoError(t, err)
	})

	return post
}
