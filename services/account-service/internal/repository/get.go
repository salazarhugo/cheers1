package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/account/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
)

func (p *accountRepository) GetAccount(
	accountID string,
) (*pb.Account, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "cheers-a275e")
	if err != nil {
		return nil, err
	}
	defer client.Close()

	doc, err := client.Collection("accounts").Doc(accountID).Get(ctx)

	account := &pb.Account{}
	err = utils.MapToProto(account, doc.Data())
	if err != nil {
		return nil, err
	}

	return account, nil
}
