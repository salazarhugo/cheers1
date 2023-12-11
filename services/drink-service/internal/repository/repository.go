package repository

import (
	"github.com/go-redis/redis/v9"
	"github.com/salazarhugo/cheers1/gen/go/cheers/drink/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"gorm.io/gorm"
	"os"
)

type Repository interface {
	CreateDrink(
		name string,
		icon string,
	) (int64, error)

	GetDrink(
		drinkID int64,
	) (*drink.Drink, error)

	UpdateDrink(
		userId string,
	) error

	ListDrink() ([]*drink.Drink, error)
}

type repository struct {
	redis   *redis.Client
	spanner *gorm.DB
}

func (r repository) UpdateDrink(userId string) error {
	//TODO implement me
	panic("implement me")
}

func NewRepository() Repository {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ENDPOINT"),
		Password: os.Getenv("DB_PASSWORD"),
		DB:       0,
	})

	return &repository{
		redis:   rdb,
		spanner: utils.GetSpanner(),
	}
}
