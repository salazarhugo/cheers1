package app

import (
	"context"
	"encoding/binary"
	_ "firebase.google.com/go/v4/auth"
	"fmt"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
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

	// Validate Username
	if len(username) < 1 {
		return nil, status.Error(codes.InvalidArgument, "invalid username")
	}

	allowedOrigins := []string{
		"https://cheers.social",
		"android:apk-key-hash:Bkcj8et21W2ji7H-tughBi15WC1Xk5NMNV9rzGfd4oI",
	}

	wconfig := &webauthn.Config{
		RPDisplayName: "Cheers",        // Display Name for your site
		RPID:          "cheers.social", // Generally the FQDN for your site
		RPOrigins:     allowedOrigins,  // The origin URLs allowed for WebAuthn requests
	}

	webAuthn, err := webauthn.New(wconfig)
	if err != nil {
		fmt.Println(err)
	}

	// Get the user
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
