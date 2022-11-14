package repository

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/account/v1"
)

type AccountRepository interface {
	CreateAccount(accountID string, post *pb.Account) (string, error)
	GetAccount(accountID string, otherAccountID string) (*pb.GetAccountResponse, error)
	UpdateAccount(accountID string, account *pb.UpdateAccountRequest) error
	DeleteAccount(id string) error

	IncrementBalance(accountID string, value int32) error
}

type accountRepository struct {
}

func NewAccountRepository() AccountRepository {
	return &accountRepository{}
}
