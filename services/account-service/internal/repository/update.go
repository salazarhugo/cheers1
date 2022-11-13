package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/account/v1"
)

func (p *accountRepository) UpdateAccount(
	accountID string,
	account *pb.UpdateAccountRequest,
) error {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "cheers-a275e")
	if err != nil {
		return err
	}
	defer client.Close()

	_, err = client.Collection("users").Doc(accountID).Set(ctx, account)
	if err != nil {
		return err
	}

	return nil
}
