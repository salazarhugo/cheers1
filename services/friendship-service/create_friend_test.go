package friendship_service

import (
	"github.com/salazarhugo/cheers1/services/friendship-service/internal/repository"
	"testing"
)

func TestCreateFriend(t *testing.T) {
	//Create a mock repository
	repo := repository.NewRepository()

	// Given
	user1 := "user1"
	user2 := "user2"

	// Test
	err := repo.CreateFriend(user1, user2)
	if err != nil {
		t.Error(err)
		return
	}

	// Clean
	err = repo.DeleteFriend(user1, user2)
	if err != nil {
		t.Error(err)
		return
	}
}
