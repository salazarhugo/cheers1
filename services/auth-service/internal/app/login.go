package app

import (
	"context"
	_ "firebase.google.com/go/v4/auth"
	"github.com/salazarhugo/cheers1/gen/go/cheers/auth/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/auth-service/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) Login(
	ctx context.Context,
	request *auth.LoginRequest,
) (*auth.LoginResponse, error) {
	idToken := request.GetIdToken()

	if idToken == "" {
		return nil, status.Error(codes.InvalidArgument, "blank parameter idToken")
	}

	app := utils.InitializeAppDefault()
	client := utils.InitializeAuthClient(app)

	// Verify Id token
	token, err := client.VerifyIDTokenAndCheckRevoked(context.Background(), idToken)
	if err != nil {
		log.Println("Invalid Authorization Token or revoked or disabled")
		return nil, status.Error(codes.PermissionDenied, "invalid token")
	}

	// Check if user exists
	user, err := service.GetUser(token.UID)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.NotFound, "can't get user")
	}

	return &auth.LoginResponse{
		User: user,
	}, nil
}
