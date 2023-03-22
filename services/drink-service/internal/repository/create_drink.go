package repository

import (
	"context"
	drink2 "github.com/salazarhugo/cheers1/gen/go/cheers/drink/v1"
	"github.com/salazarhugo/cheers1/libs/utils/mapper"
)

func (r repository) CreateDrink(
	creatorID string,
	name string,
	icon string,
	category string,
) (string, error) {
	ctx := context.Background()
	doc := r.firestore.Collection("drinks").NewDoc()
	drink := &drink2.Drink{
		Id:        doc.ID,
		CreatorId: creatorID,
		Name:      name,
		Icon:      icon,
		Category:  category,
	}

	m, err := mapper.ProtoToMap(drink)
	_, err = doc.Set(ctx, m)
	if err != nil {
		return "", err
	}

	return doc.ID, nil
}
