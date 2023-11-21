package app

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	_ "firebase.google.com/go/v4/auth"
	"fmt"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/salazarhugo/cheers1/gen/go/cheers/auth/v1"
	"github.com/salazarhugo/cheers1/services/auth-service/internal/repository"
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

	user, _ := s.authRepository.GetUserByUsername(request.GetUsername())

	var authnUser *repository.AuthnUser
	if user != nil {
		authnUser = repository.NewUser(
			user.AuthnId,
			username,
			username,
		)
	} else {
		authnUser = repository.NewUser(
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

	num := binary.LittleEndian.Uint64(sessionData.UserID)
	log.Println(uint64(sessionData.UserID))
	log.Println(sessionData.Challenge)

	return &auth.BeginRegistrationResponse{
		UserId:    string(sessionData.UserID),
		Challenge: sessionData.Challenge,
	}, nil
}

func randomUint64() uint64 {
	buf := make([]byte, 8)
	rand.Read(buf)
	return binary.LittleEndian.Uint64(buf)
}
