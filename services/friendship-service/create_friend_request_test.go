package friendship_service

import (
	"github.com/salazarhugo/cheers1/services/friendship-service/internal/repository"
	"log"
	"testing"
)

func TestCreateFriendRequest(t *testing.T) {
	//Create a mock repository
	repo := repository.NewRepository()

	// Given
	user1 := "user1"
	user2 := "user2"

	// Test
	err := repo.CreateFriendRequest(user1, user2)
	if err != nil {
		t.Error(err)
		return
	}

	friendRequests, err := repo.ListFriendRequests(0, 10, user2)
	if err != nil {
		t.Error(err)
		return
	}

	log.Println(friendRequests)

	err = repo.DeleteFriendRequest(user1, user2)
	if err != nil {
		t.Error(err)
		return
	}
}
