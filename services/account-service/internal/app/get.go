package app

import (
	"context"
	"github.com/labstack/gommon/log"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/account/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetAccount(
	ctx context.Context,
	request *pb.GetAccountRequest,
) (*pb.GetAccountResponse, error) {
	accountID, err := GetAccountId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed retrieving accountID")
	}

	otherAccountID := request.GetAccountId()

	response, err := s.accountRepository.GetAccount(accountID, otherAccountID)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, "failed to get account")
	}

	return response, nil
}
