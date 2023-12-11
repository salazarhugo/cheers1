package user_service

import (
	"github.com/salazarhugo/cheers1/services/drink-service/internal/repository"
	"log"
	"testing"
)

func TestListDrinks(t *testing.T) {
	repo := repository.NewRepository()

	drink, err := repo.ListDrink()
	if err != nil {
		return
	}
	log.Println(drink)
}
