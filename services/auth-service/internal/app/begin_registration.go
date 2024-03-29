package app

import (
	"context"
	"encoding/binary"
	_ "firebase.google.com/go/v4/auth"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/salazarhugo/cheers1/gen/go/cheers/auth/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) BeginRegistration(
	ctx context.Context,
	request *auth.BeginRegistrationRequest,
) (*auth.BeginRegistrationResponse, error) {
	username := request.GetUsername()

	// Validate username
	isTaken, err := s.authRepository.CheckUsername(username)
	if err != nil {
		return nil, err
	}
	if isTaken == true {
		return nil, status.Error(codes.AlreadyExists, "username already exists")
	}

	webAuthn, err := NewWebAuthn()
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to initialize webAuthn")
	}

	// Create the user
	authnUser, err := s.authRepository.GetOrCreateUser(username)
	if err != nil {
		return nil, err
	}

	registerOptions := func(credCreationOpts *protocol.PublicKeyCredentialCreationOptions) {
		credCreationOpts.CredentialExcludeList = authnUser.CredentialExcludeList()
	}

	_, sessionData, err := webAuthn.BeginRegistration(
		authnUser,
		registerOptions,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	num := binary.LittleEndian.Uint64(sessionData.UserID)

	return &auth.BeginRegistrationResponse{
		UserId:    num,
		Challenge: sessionData.Challenge,
	}, nil
}
