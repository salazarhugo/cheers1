package auth_service

import (
	"github.com/salazarhugo/cheers1/services/auth-service/internal/repository"
	"github.com/salazarhugo/cheers1/services/auth-service/internal/service"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	//Create a mock repository
	repo := repository.NewRepository()

	user, err := service.CreateUser(
		"hugorb@cheers.social",
		"donald",
	)
	if err != nil {
		t.Error("failed to create user: ", err)
		return
	}

	_, err = repo.CreateCredential(
		user.Id,
		"publickey23r52582",
	)
	if err != nil {
		t.Error("failed to create credential: ", err)
		return
	}
}
