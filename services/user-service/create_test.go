package user_service

import (
	"github.com/salazarhugo/cheers1/services/user-service/internal/repository"
	"testing"
)

func TestCreateUser(t *testing.T) {
	// Create a mock repository
	repo := repository.NewUserRepository()

	_, err := repo.CreateUser(
		"user-01",
		"nike2",
		"Nike",
		"picture",
		"hugobrock74+nike@gmail.com",
	)

	if err != nil {
		t.Error("Failed to create user: ", err)
	}
}
