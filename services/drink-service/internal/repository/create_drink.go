package repository

import (
	"github.com/google/uuid"
	"github.com/salazarhugo/cheers1/services/drink-service/internal/domain"
)

func (r repository) CreateDrink(
	name string,
	icon string,
) (int64, error) {
	db := r.spanner

	drink := &domain.Drink{
		ID:   int64(uuid.New().ID()),
		Name: name,
		Icon: icon,
	}

	result := db.Create(&drink)
	if result.Error != nil {
		return -1, result.Error
	}

	return drink.ID, nil
}
