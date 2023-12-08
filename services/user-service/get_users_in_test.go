package user_service

import (
	"github.com/salazarhugo/cheers1/services/user-service/internal/repository"
	"log"
	"testing"
)

func TestGetUsersIn(t *testing.T) {
	//Create a mock repository
	repo := repository.NewUserRepository()

	query := []string{"dream", "cheers"}

	users, err := repo.GetUsersIn(query)
	if err != nil {
		t.Error("failed to check username: ", err)
		return
	}

	for _, user := range users {
		log.Println(user.Username)
	}
}
