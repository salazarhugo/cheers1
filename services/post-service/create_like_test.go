package user_service

import (
	"cloud.google.com/go/spanner"
	"github.com/salazarhugo/cheers1/services/post-service/internal/repository"
	"testing"
)

func TestCreateLike(t *testing.T) {
	//Create a mock repository
	repo := repository.NewPostRepository()
	userID := "user1"

	caption := "Only dead fish go with the flow"
	postID, err := repo.CreatePost(
		userID,
		&repository.Post{
			Caption: caption,
			DrinkID: spanner.NullInt64{Int64: -1, Valid: false},
			Photos:  []string{"wfwf", "wef"},
		},
	)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = repo.LikePost(
		userID,
		postID,
	)
	if err != nil {
		t.Error(err)
	}

	_, err = repo.UnlikePost(
		userID,
		postID,
	)
	if err != nil {
		t.Error(err)
	}

	err = repo.DeletePostById(postID)
	if err != nil {
		t.Error(err)
		return
	}
}
