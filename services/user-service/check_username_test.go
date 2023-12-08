package user_service

import (
	"github.com/salazarhugo/cheers1/services/user-service/internal/repository"
	"log"
	"testing"
)

func TestCheckUsername(t *testing.T) {
	//Create a mock repository
	repo := repository.NewUserRepository()

	username := "hugo"
	exists, err := repo.CheckUsername(username)

	if err != nil {
		t.Error("failed to check username: ", err)
		return
	}

	log.Println(exists)
}
