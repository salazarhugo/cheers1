package user_service

import (
	"cloud.google.com/go/spanner"
	"github.com/salazarhugo/cheers1/services/post-service/internal/repository"
	"testing"
)

func TestCreatePost(t *testing.T) {
	//Create a mock repository
	repo := repository.NewPostRepository()

	caption := "Only dead fish go with the flow"
	userID, err := repo.CreatePost(
		"user1",
		&repository.Post{
			Caption: caption,
			DrinkID: spanner.NullInt64{Int64: 1, Valid: true},
			Photos:  []string{"wfwf", "wef"},
		},
	)

	if err != nil {
		t.Error("failed to create user: ", err)
		return
	}

	post, err := repo.GetPostById(userID)
	if err != nil {
		t.Error("failed to get user: ", err)
		return
	}
	if post.Caption != caption {
		t.Error("wrong post caption: ", err)
		return
	}

	err = repo.DeletePost(userID)
	if err != nil {
		t.Error("failed to delete user: ", err)
		return
	}
}
