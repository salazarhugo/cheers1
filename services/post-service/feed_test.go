package user_service

import (
	"github.com/salazarhugo/cheers1/services/post-service/internal/repository"
	"log"
	"testing"
)

func TestFeedPost(t *testing.T) {
	//Create a mock repository
	repo := repository.NewPostRepository()

	friendIDs := []string{"user1", "cheers"}
	response, err := repo.FeedPost(
		friendIDs,
		1,
		10,
	)

	if err != nil {
		t.Error("failed to get feed: ", err)
		return
	}

	log.Println(response)
}
