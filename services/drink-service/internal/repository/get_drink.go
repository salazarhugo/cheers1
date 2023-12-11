package repository

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/drink/v1"
	"log"
)

func (r repository) GetDrink(
	drinkID int64,
) (*drink.Drink, error) {
	db := r.spanner
	log.Println(db)
	return nil, nil
}
