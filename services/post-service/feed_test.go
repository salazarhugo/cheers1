package user_service

import (
	"github.com/salazarhugo/cheers1/services/post-service/internal/repository"
	"log"
	"testing"
)

func TestFeedPost(t *testing.T) {
	//Create a mock repository
	repo := repository.NewPostRepository()

	friendIDs := []string{"Ii7mFI34JaYKYsOxYg90fYToa182", "a082acf0-d046-4467-852e-9eeee17df56f"}
	response, err := repo.FeedPost(
		friendIDs,
		1,
		10,
	)

	if err != nil {
		t.Error("failed to get feed: ", err)
		return
	}

	for _, postResponse := range response.GetPosts() {
		log.Println(postResponse)
	}
}
