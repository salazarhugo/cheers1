package domain

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/drink/v1"
	"time"
)

// Drink model
type Drink struct {
	ID        int64 `gorm:"primarykey"`
	Name      string
	Icon      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (d Drink) ToDrinkPb() *drink.Drink {
	return &drink.Drink{
		Id:        d.ID,
		CreatorId: "",
		Name:      d.Name,
		Icon:      d.Icon,
		Category:  "",
	}
}
