package app

import (
	"context"
	"github.com/labstack/gommon/log"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/account/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteAccount(
	ctx context.Context,
	request *pb.DeleteAccountRequest,
) (*pb.DeleteAccountResponse, error) {
	_, err := GetAccountId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving accountID")
	}

	id := request.GetAccountId()
	if id == "" {
		return nil, status.Error(codes.InvalidArgument, "id parameter can't be empty")
	}

	err = s.accountRepository.DeleteAccount(id)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, "failed delete account")
	}

	return &pb.DeleteAccountResponse{}, nil
}
