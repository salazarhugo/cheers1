package repository

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/account/v1"
)

func (p *accountRepository) GetAccount(
	accountID string,
	otherAccountID string,
) (*pb.GetAccountResponse, error) {
	return nil, nil
}
