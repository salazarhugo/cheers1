package app

import (
	"context"
	"encoding/binary"
	_ "firebase.google.com/go/v4/auth"
	"fmt"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/salazarhugo/cheers1/gen/go/cheers/auth/v1"
	"log"
)

func (s *Server) BeginLogin(
	ctx context.Context,
	request *auth.BeginLoginRequest,
) (*auth.BeginLoginResponse, error) {
	username := request.GetUsername()

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
	authnUser, err := s.authRepository.GetAuthnUser(username)
	if err != nil {
		return nil, err
	}

	// Generate PublicKeyCredentialRequestOptions
	options, sessionData, err := webAuthn.BeginLogin(authnUser)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	num := binary.LittleEndian.Uint64(sessionData.UserID)

	return &auth.BeginLoginResponse{
		UserId:           num,
		Challenge:        sessionData.Challenge,
		RelyingPartyId:   options.Response.RelyingPartyID,
		UserVerification: string(options.Response.UserVerification),
		AllowCredentials: authnUser.WebAuthnCredentials(),
		Timeout:          60 * 1000, // in milliseconds
	}, nil
}
