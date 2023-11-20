package app

import (
	"context"
	_ "firebase.google.com/go/v4/auth"
	"fmt"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/salazarhugo/cheers1/gen/go/cheers/auth/v1"
	"log"
)

func (s *Server) BeginRegistration(
	ctx context.Context,
	request *auth.BeginRegistrationRequest,
) (*auth.BeginRegistrationResponse, error) {
	username := request.GetUsername()
	wconfig := &webauthn.Config{
		RPDisplayName: "Cheers",                          // Display Name for your site
		RPID:          "go-webauthn.local",               // Generally the FQDN for your site
		RPOrigins:     []string{"https://cheers.social"}, // The origin URLs allowed for WebAuthn requests
	}

	webAuthn, err := webauthn.New(wconfig)
	if err != nil {
		fmt.Println(err)
	}

	user, err := s.authRepository.GetUserByUsername(request.GetUsername())

	var authnUser *AuthnUser
	if user != nil {
		authnUser = NewUser(
			user.AuthnId,
			username,
			username,
		)
	} else {
		authnUser = NewUser(
			randomUint64(),
			username,
			username,
		)
	}

	_, sessionData, err := webAuthn.BeginRegistration(
		authnUser,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = s.authRepository.PutSession(username, sessionData)
	if err != nil {
		return nil, err
	}

	return &auth.BeginRegistrationResponse{
		UserId:    string(sessionData.UserID),
		Challenge: sessionData.Challenge,
	}, nil
}
