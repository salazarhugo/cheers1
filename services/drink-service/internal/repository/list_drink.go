package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/drink/v1"
	"github.com/salazarhugo/cheers1/libs/utils/mapper"
	"google.golang.org/api/iterator"
)

func (r repository) ListDrink() ([]*drink.Drink, error) {
	ctx := context.Background()
	documents := r.firestore.Collection("drinks").OrderBy("name", firestore.Asc).Documents(ctx)

	drinks := make([]*drink.Drink, 0)

	for {
		doc, err := documents.Next()
		if err == iterator.Done {
			break
		}
		item := &drink.Drink{}
		m := doc.Data()
		err = mapper.MapToProto(item, m)
		drinks = append(drinks, item)
	}

	return drinks, nil
}
