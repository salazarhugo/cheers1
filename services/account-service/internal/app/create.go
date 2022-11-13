package app

import (
	"context"
	"github.com/labstack/gommon/log"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/account/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAccount(
	ctx context.Context,
	request *pb.CreateAccountRequest,
) (*pb.CreateAccountResponse, error) {
	accountID, err := GetAccountId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving accountID")
	}

	partyReq := request.GetAccount()
	if partyReq == nil {
		return nil, status.Error(codes.InvalidArgument, "post parameter can't be nil")
	}

	accountID, err = s.accountRepository.CreateAccount(accountID, partyReq)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, "failed to create account")
	}

	account, err := s.accountRepository.GetAccount(accountID, accountID)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, "failed to get account")
	}

	return &pb.CreateAccountResponse{
		Account: account.Account,
	}, nil
}
