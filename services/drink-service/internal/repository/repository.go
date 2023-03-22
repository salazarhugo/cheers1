package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/salazarhugo/cheers1/gen/go/cheers/drink/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"os"
)

type Repository interface {
	CreateDrink(
		creatorID string,
		name string,
		icon string,
		category string,
	) error

	UpdateDrink(
		userId string,
		latitude float64,
		longitude float64,
	) error

	ListDrink(
		userId string,
	) ([]*drink.Drink, error)
}

type repository struct {
	redis     *redis.Client
	firestore *firestore.Client
}

func NewRepository() Repository {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ENDPOINT"),
		Password: os.Getenv("DB_PASSWORD"),
		DB:       0,
	})
	fire, err := utils.InitializeAppDefault().Firestore(context.Background())
	if err != nil {
		return nil
	}

	return &repository{
		redis:     rdb,
		firestore: fire,
	}
}
