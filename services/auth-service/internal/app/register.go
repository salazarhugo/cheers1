package app

import (
	"context"
	_ "firebase.google.com/go/v4/auth"
	"github.com/salazarhugo/cheers1/gen/go/cheers/auth/v1"
	"github.com/salazarhugo/cheers1/services/auth-service/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Register(
	ctx context.Context,
	request *auth.RegisterRequest,
) (*auth.RegisterResponse, error) {

	// Validate Email and Username
	if len(request.Email) < 1 || len(request.Username) < 1 {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}

	user, err := service.CreateUser(
		request.Email,
		request.Username,
	)
	if err != nil {
		return nil, err
	}

	_, err = s.authRepository.CreateCredential(
		request.PublicKey,
		user.Id,
	)
	if err != nil {
		return nil, err
	}

	return &auth.RegisterResponse{
		User: user,
	}, nil
}
