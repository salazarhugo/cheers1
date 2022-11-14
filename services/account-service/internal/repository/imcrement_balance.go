package repository

import (
	"cloud.google.com/go/firestore"
	"context"
)

func (p *accountRepository) IncrementBalance(
	accountID string,
	value int32,
) error {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "cheers-a275e")
	if err != nil {
		return err
	}
	defer client.Close()

	_, err = client.Collection("accounts").Doc(accountID).Set(ctx,
		map[string]interface{}{
			"balance": firestore.Increment(value),
		},
	)
	if err != nil {
		return err
	}

	return nil
}
