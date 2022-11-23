package app

import (
	"context"
	"github.com/labstack/gommon/log"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/account/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetAccount(
	ctx context.Context,
	request *pb.GetAccountRequest,
) (*pb.GetAccountResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed retrieving userID")
	}

	response, err := s.accountRepository.GetAccount(userID)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, "failed to get account")
	}

	return &pb.GetAccountResponse{
		Account: response,
	}, nil
}
