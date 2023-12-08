package auth_service

import (
	"github.com/salazarhugo/cheers1/services/auth-service/internal/repository"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestCheckUsernameExists(t *testing.T) {
	//Create a mock repository
	repo := repository.NewRepository()

	username := "cheers"
	exists, err := repo.CheckUsername(username)

	if err != nil {
		t.Error("failed to check username: ", err)
		return
	}
	assert.Equal(t, true, exists)
}

func TestCheckUsernameNotExist(t *testing.T) {
	//Create a mock repository
	repo := repository.NewRepository()

	username := "hfjekhgertetekjghekrghkerjhg"
	exists, err := repo.CheckUsername(username)

	if err != nil {
		t.Error("failed to check username: ", err)
		return
	}
	assert.Equal(t, false, exists)

	log.Println(exists)
}
