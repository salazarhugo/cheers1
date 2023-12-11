package repository

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/drink/v1"
	"github.com/salazarhugo/cheers1/services/drink-service/internal/domain"
)

func (r repository) ListDrink() ([]*drink.Drink, error) {
	db := r.spanner

	var drinks []*domain.Drink

	result := db.Table("drinks").Order("name asc").Find(&drinks)
	if result.Error != nil {
		return nil, result.Error
	}

	items := make([]*drink.Drink, 0)
	for _, drink := range drinks {
		items = append(items, drink.ToDrinkPb())
	}

	return items, nil
}
