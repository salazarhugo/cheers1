package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/auth/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateBusinessAccount(
	ctx context.Context,
	request *auth.CreateBusinessAccountRequest,
) (*auth.CreateBusinessAccountResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed retrieving userID")
	}

	app := utils.InitializeAppDefault()
	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	// Check if user is admin
	user, err := client.GetUser(ctx, userID)
	if user.CustomClaims["admin"] == nil {
		return nil, status.Error(codes.PermissionDenied, "insufficient permissions")
	}

	// Create Business Account
	err = s.authRepository.CreateBusinessAccount(request.UserId)
	if err != nil {
		return nil, err
	}

	return &auth.CreateBusinessAccountResponse{}, nil
}
