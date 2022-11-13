package app

import (
	"context"
	"github.com/labstack/gommon/log"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/account/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateAccount(
	ctx context.Context,
	request *pb.UpdateAccountRequest,
) (*pb.UpdateAccountResponse, error) {
	accountID, err := GetAccountId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed retrieving accountID")
	}

	err = s.accountRepository.UpdateAccount(accountID, request)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, "failed to update account")
	}

	result, err := s.accountRepository.GetAccount(accountID, accountID)

	return &pb.UpdateAccountResponse{
		Account: result.GetAccount(),
	}, nil
}
