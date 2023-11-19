package app

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/binary"
	_ "firebase.google.com/go/v4/auth"
	"fmt"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/salazarhugo/cheers1/gen/go/cheers/auth/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"log"
	"net/http"
)

// User represents the user model
type User struct {
	id          uint64
	name        string
	displayName string
	credentials []webauthn.Credential
}

func (u User) WebAuthnID() []byte {
	return []byte{}
}

func (u User) WebAuthnName() string {
	//TODO implement me
	panic("implement me")
}

func (u User) WebAuthnDisplayName() string {
	//TODO implement me
	panic("implement me")
}

func (u User) WebAuthnCredentials() []webauthn.Credential {
	//TODO implement me
	panic("implement me")
}

func (u User) WebAuthnIcon() string {
	//TODO implement me
	panic("implement me")
}

// NewUser creates and returns a new User
func NewUser(name string, displayName string) *User {

	user := &User{}
	user.id = randomUint64()
	user.name = name
	user.displayName = displayName
	// user.credentials = []webauthn.Credential{}

	return user
}

func randomUint64() uint64 {
	buf := make([]byte, 8)
	rand.Read(buf)
	return binary.LittleEndian.Uint64(buf)
}

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
	passkeyBytes, err := proto.Marshal(request.Passkey)
	if err != nil {
		log.Fatal("Error marshaling protobuf message:", err)
	}

	bodyReader := bytes.NewBuffer(passkeyBytes)
	req, err := http.NewRequest(http.MethodGet, "", bodyReader)
	if err != nil {
		return nil, err
	}

	user := NewUser(request.Username, request.Username)
	log.Println(user)
	log.Println(req)

	// generate PublicKeyCredentialCreationOptions, session data
	_, sessionData, err := webAuthn.BeginRegistration(
		user,
	)

	credential, err := webAuthn.FinishRegistration(user, *sessionData, req)
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
