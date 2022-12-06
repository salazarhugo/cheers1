package repository

import (
	"cloud.google.com/go/firestore"
	"context"
)

func (p *accountRepository) IncrementBalance(
	accountID string,
	value int64,
) error {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "cheers-a275e")
	if err != nil {
		return err
	}
	defer client.Close()

	doc, err := client.Collection("accounts").Doc(accountID).Get(ctx)

	if doc != nil && doc.Exists() {
		_, err = client.Collection("accounts").Doc(accountID).Update(ctx,
			[]firestore.Update{
				{Path: "balance", Value: firestore.Increment(value)},
			},
		)
	} else {
		_, err = client.Collection("accounts").Doc(accountID).Set(ctx,
			map[string]interface{}{
				"balance": value,
			},
		)
	}
	if err != nil {
		return err
	}

	return nil
}
