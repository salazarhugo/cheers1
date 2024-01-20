package user_service

import (
	"github.com/salazarhugo/cheers1/services/user-service/internal/repository"
	"testing"
)

func TestVerifyUser(t *testing.T) {
	//Create a mock repository
	repo := repository.NewUserRepository()

	userID := "Ii7mFI34JaYKYsOxYg90fYToa182"

	err := repo.VerifyUser(
		userID,
	)

	if err != nil {
		t.Error(err)
		return
	}
}
