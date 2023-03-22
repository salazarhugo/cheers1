package repository

import (
	"context"
	"fmt"
)

const (
	keyUser        = "user"
	keyUserDrinks  = "drinks"
	keyLastUpdated = "updated"
)

func getKeyUserDrinks(userId string) string {
	return fmt.Sprintf("%s:%s", userId, userId)
}

func (r repository) UpdateDrink(
	userId string,
	latitude float64,
	longitude float64,
) error {
	_ = context.Background()

	return nil
}
