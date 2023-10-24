package auth

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CheckIsAdmin Check if user is admin
func CheckIsAdmin(userID string) (bool, error) {
	ctx := context.Background()
	app := InitializeAppDefault()
	client, err := app.Auth(ctx)
	if err != nil {
		return false, err
	}

	user, err := client.GetUser(ctx, userID)
	if user.CustomClaims["admin"] == nil {
		return false, status.Error(codes.PermissionDenied, "insufficient permissions: not admin")
	}

	return true, nil
}
