package http

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/notification/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) CreateRegistrationToken(
	ctx context.Context,
	request *notification.CreateRegistrationTokenRequest,
) (*notification.CreateRegistrationTokenResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed retrieving userID")
	}

	token := request.GetToken()
	if token == "" {
		return nil, status.Error(codes.InvalidArgument, "registrationToken parameter can't be nil")
	}

	err = s.repository.CreateRegistrationToken(userID, token)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &notification.CreateRegistrationTokenResponse{}, nil
}
