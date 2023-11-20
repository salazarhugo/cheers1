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

func (s *Server) FinishRegistration(
	ctx context.Context,
	request *auth.FinishRegistrationRequest,
) (*auth.FinishRegistrationResponse, error) {
	log.Println(request)

	wconfig := &webauthn.Config{
		RPDisplayName: "Cheers",                          // Display Name for your site
		RPID:          "go-webauthn.local",               // Generally the FQDN for your site
		RPOrigins:     []string{"https://cheers.social"}, // The origin URLs allowed for WebAuthn requests
	}

	webAuthn, err := webauthn.New(wconfig)
	if err != nil {
		fmt.Println(err)
	}

	// Marshal the protobuf message into a byte slice
	passkeyBytes, err := protojson.Marshal(request.Passkey)
	if err != nil {
		log.Fatal("Error marshaling protobuf message:", err)
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

	if err != nil {
		log.Println(err)
		return nil, err
	}

	sessionData := webauthn.SessionData{
		Challenge:            request.Challenge,
		UserID:               []byte(request.UserId),
		AllowedCredentialIDs: nil,
		Expires:              time.Now().AddDate(0, 0, 1),
		UserVerification:     "",
		Extensions:           nil,
	}
	if err != nil {
		log.Fatal("failed to get session: ", err)
		return nil, err
	}

	credential, err := webAuthn.FinishRegistration(user.ToAuthnUser(), sessionData, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(credential)
	log.Println("SUCCESS")

	// Validate Email and Username
	if len(request.Email) < 1 || len(request.Username) < 1 {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}

	//user, err := service.CreateUser(
	//	request.Email,
	//	request.Username,
	//)
	//if err != nil {
	//	return nil, err
	//}
	//
	//_, err = s.authRepository.CreateCredential(
	//	request.PublicKey,
	//	user.Id,
	//)
	//if err != nil {
	//	return nil, err
	//}

	return &auth.FinishRegistrationResponse{
		User: nil,
	}, nil
}
