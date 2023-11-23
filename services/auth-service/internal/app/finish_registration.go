package app

import (
	"bytes"
	"context"
	_ "firebase.google.com/go/v4/auth"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/salazarhugo/cheers1/gen/go/cheers/auth/v1"
	"github.com/salazarhugo/cheers1/services/auth-service/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net/http"
	"time"
)

func (s *Server) FinishRegistration(
	ctx context.Context,
	request *auth.FinishRegistrationRequest,
) (*auth.FinishRegistrationResponse, error) {
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

	user, err := s.authRepository.GetUserByUsername(request.Username)
	if err != nil {
		return nil, err
	}

	authnUser := user.ToAuthnUser()

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

	credential, err := webAuthn.FinishRegistration(authnUser, sessionData, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	authnUser.AddCredential(*credential)
	log.Println(credential.PublicKey)

	err = s.authRepository.CreateCredential(repository.CredentialToDomain(user.AuthnId, credential))
	if err != nil {
		return nil, err
	}

	return &auth.FinishRegistrationResponse{
		User: user.ToUserPb(),
	}, nil
}
