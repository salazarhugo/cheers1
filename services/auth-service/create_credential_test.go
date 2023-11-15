package auth_service

import (
	"github.com/salazarhugo/cheers1/services/auth-service/internal/repository"
	"testing"
)

func TestCreateCredential(t *testing.T) {
	//Create a mock repository
	repo := repository.NewRepository()

	_, err := repo.CreateCredential(
		"user1",
		"publickey23r52582",
	)
	if err != nil {
		t.Error("failed to create credential: ", err)
		return
	}
}
