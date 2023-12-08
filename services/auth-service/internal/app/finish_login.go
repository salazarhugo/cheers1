package app

import (
	"bytes"
	"context"
	_ "firebase.google.com/go/v4/auth"
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
	// Initialize webAuthn
	webAuthn, err := NewWebAuthn()
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to initialize webAuthn")
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

	// FIDO validate login
	_, err = webAuthn.FinishLogin(authnUser, sessionData, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Get the user
	user, err := s.authRepository.GetUserByUsername(request.Username)
	if err != nil {
		return nil, err
	}

	// Create secure JSON web token (JWT)
	token, err := s.authRepository.CreateFirebaseCustomToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &auth.FinishLoginResponse{
		User:  user.ToUserPb(),
		Token: token,
	}, nil
}
