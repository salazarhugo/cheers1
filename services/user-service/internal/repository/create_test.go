package repository

import (
	"testing"
)

func TestCreateUser(t *testing.T) {
	// Create a mock repository
	repo := NewUserRepository()

	_, err := repo.CreateUser(
		"user-01",
		"nike",
		"Nike",
		"picture",
		"hugobrock74+nike@gmail.com",
	)

	if err != nil {
		t.Error("Failed to create user")
	}
}
