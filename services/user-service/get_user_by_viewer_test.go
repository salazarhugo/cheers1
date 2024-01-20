package user_service

import (
	"github.com/salazarhugo/cheers1/services/user-service/internal/repository"
	"log"
	"testing"
)

func TestGetUserByViewer(t *testing.T) {
	//Create a mock repository
	repo := repository.NewUserRepository()

	// GIVEN
	viewerID := "Hrpo40ykFThEkIoFHercdmqGO0y2"
	userID := "Hrpo40ykFThEkIoFHercdmqGO0y2"

	user, err := repo.GetUserWithViewer(viewerID, userID)
	if err != nil {
		t.Error(err)
		return
	}

	log.Println(user)
}
