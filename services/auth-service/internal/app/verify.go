package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/auth/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) VerifyUser(
	ctx context.Context,
	request *auth.VerifyUserRequest,
) (*auth.VerifyUserResponse, error) {
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

	// Set business account privilege on the user corresponding to uid.
	claims := user.CustomClaims
	claims["verified"] = true
	err = client.SetCustomUserClaims(ctx, request.GetUserId(), claims)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "error setting custom claims")
	}

	return &auth.VerifyUserResponse{}, nil
}
