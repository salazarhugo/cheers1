package user_service

import (
	"github.com/salazarhugo/cheers1/services/post-service/internal/repository"
	"log"
	"testing"
)

func TestFeedPost(t *testing.T) {
	//Create a mock repository
	repo := repository.NewPostRepository()

	friendIDs := []string{"cheers", "user1"}
	response, err := repo.FeedPost(
		friendIDs,
		5,
		10,
	)

	if err != nil {
		t.Error("failed to get feed: ", err)
		return
	}

	for _, postResponse := range response.GetPosts() {
		log.Println(postResponse.HasLiked)
		log.Println(postResponse.LikeCount)
	}
}
