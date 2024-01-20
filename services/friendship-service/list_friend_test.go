package friendship_service

import (
	"github.com/salazarhugo/cheers1/services/friendship-service/internal/repository"
	"log"
	"testing"
)

func TestListFriend(t *testing.T) {
	//Create a mock repository
	repo := repository.NewRepository()

	// Given
	viewerID := "user2"
	user2 := "user2"

	// Test
	users, err := repo.ListFriend(viewerID, 1, 100, user2)
	if err != nil {
		t.Error(err)
		return
	}

	log.Println(users)
}
