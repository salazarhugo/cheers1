package user_service

import (
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
			DrinkID: 1,
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

	err = repo.DeletePostById(userID)
	if err != nil {
		t.Error("failed to delete user: ", err)
		return
	}
}
