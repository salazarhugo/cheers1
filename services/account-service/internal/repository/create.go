package repository

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/account/v1"
)

func (p *accountRepository) CreateAccount(
	accountID string,
	account *pb.Account,
) (string, error) {
	return account.Id, nil
}
