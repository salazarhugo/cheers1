package user_service

import (
	"github.com/salazarhugo/cheers1/services/user-service/internal/repository"
	"testing"
)

func TestDeleteUser(t *testing.T) {
	// Create a mock repository
	repo := repository.NewUserRepository()

	userID := "user-01"
	_, err := repo.CreateUser(
		userID,
		"nike2",
		"Nike",
		"picture",
		"hugobrock74+nike@gmail.com",
	)

	if err != nil {
		t.Error("failed to create user: ", err)
	}

	err = repo.DeleteUser(userID)
	if err != nil {
		t.Error("failed to delete user: ", err)
	} else {
		t.Logf("Successfully deleted user: %s", userID)
	}

	user, err := repo.GetUserNode(userID)
	if err != nil {
		t.Error("failed to get user: ", err)
	}

	t.Log(user)
}
