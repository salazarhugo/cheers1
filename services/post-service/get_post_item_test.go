package user_service

import (
	"github.com/salazarhugo/cheers1/services/post-service/internal/repository"
	"log"
	"testing"
)

func TestGetPostItem(t *testing.T) {
	//Create a mock repository
	repo := repository.NewPostRepository()

	response, err := repo.GetPostItem(
		"user1",
		"f44d9ffb-f892-4332-bbe5-b3ca0cfae456",
	)
	if err != nil {
		t.Error("failed to get feed: ", err)
		return
	}

	log.Println(response)
}
