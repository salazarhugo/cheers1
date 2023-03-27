package repository

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/drink/v1"
	"github.com/salazarhugo/cheers1/libs/utils/mapper"
)

func (r repository) GetDrink(
	drinkID string,
) (*drink.Drink, error) {
	ctx := context.Background()
	doc, err := r.firestore.Collection("drinks").Doc(drinkID).Get(ctx)
	if err != nil {
		return nil, err
	}

	res := &drink.Drink{}
	err = mapper.MapToProto(res, doc.Data())
	if err != nil {
		return nil, err
	}

	return res, nil
}
