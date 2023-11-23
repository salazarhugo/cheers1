package auth_service

import (
	"github.com/salazarhugo/cheers1/services/auth-service/internal/repository"
	"log"
	"testing"
)

func TestGetCredentials(t *testing.T) {
	//Create a mock repository
	repo := repository.NewRepository()

	credentials, err := repo.GetUserCredentials("lars")
	if err != nil {
		t.Error("failed to create credential: ", err)
		return
	}

	user := repository.NewUser(45, "user", "wr")
	for _, credential := range credentials {
		user.AddCredential(*credential.ToWebAuthnCredential())
	}

	log.Println(user.WebAuthnCredentials())
}
