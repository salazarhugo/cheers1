package app

import (
	"bytes"
	"context"
	_ "firebase.google.com/go/v4/auth"
	"fmt"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/salazarhugo/cheers1/gen/go/cheers/auth/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net/http"
	"time"
)

func (s *Server) FinishLogin(
	ctx context.Context,
	request *auth.FinishLoginRequest,
) (*auth.FinishLoginResponse, error) {
	log.Println(request)

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

	// Marshal the protobuf message into a byte slice
	passkeyBytes, err := protojson.Marshal(request.Passkey)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	bodyReader := bytes.NewBuffer(passkeyBytes)
	req, err := http.NewRequest(http.MethodPost, "", bodyReader)
	if err != nil {
		return nil, err
	}

	authnUser, err := s.authRepository.GetAuthnUser(request.Username)
	if err != nil {
		return nil, err
	}

	sessionData := webauthn.SessionData{
		Challenge:            request.Challenge,
		UserID:               authnUser.WebAuthnID(),
		AllowedCredentialIDs: nil,
		Expires:              time.Now().AddDate(0, 0, 1),
		UserVerification:     "",
		Extensions:           nil,
	}
	if err != nil {
		log.Fatal("failed to get session: ", err)
		return nil, err
	}

	credential, err := webAuthn.FinishLogin(authnUser, sessionData, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(credential.PublicKey)

	return &auth.FinishLoginResponse{
		User: nil,
	}, nil
}
