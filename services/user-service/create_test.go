package user_service

import (
	"github.com/salazarhugo/cheers1/services/user-service/internal/repository"
	"testing"
)

func TestCreateUser(t *testing.T) {
	//Create a mock repository
	repo := repository.NewUserRepository()

	userName := "User Service"
	userID, err := repo.CreateUser(
		"user3",
		"test3",
		userName,
		"",
		"test2@cheers.social",
	)

	if err != nil {
		t.Error("failed to create user: ", err)
		return
	}

	user, err := repo.GetUserById(userID)
	if err != nil {
		t.Error("failed to get user: ", err)
		return
	}
	if user.Name != userName {
		t.Error("wrong user: ", err)
		return
	}

	err = repo.DeleteUserById(userID)
	if err != nil {
		t.Error("failed to delete user: ", err)
		return
	}
}
