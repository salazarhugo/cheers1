package app

import (
	"crypto/rand"
	"encoding/binary"
	"github.com/go-webauthn/webauthn/webauthn"
)

// User represents the user model
type User struct {
	id          uint64
	name        string
	displayName string
	credentials []webauthn.Credential
}

// WebAuthnID returns the user's ID
func (u User) WebAuthnID() []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	binary.PutUvarint(buf, uint64(u.id))
	return buf
}

func (u User) WebAuthnName() string {
	return u.name
}

func (u User) WebAuthnDisplayName() string {
	return u.displayName
}

func (u User) WebAuthnCredentials() []webauthn.Credential {
	return u.credentials
}

// AddCredential associates the credential to the user
func (u User) AddCredential(cred webauthn.Credential) {
	u.credentials = append(u.credentials, cred)
}

func (u User) WebAuthnIcon() string {
	return ""
}

// NewUser creates and returns a new User
func NewUser(
	name string,
	displayName string,
) *User {
	user := &User{}
	user.id = randomUint64()
	user.name = name
	user.displayName = displayName

	return user
}

func randomUint64() uint64 {
	buf := make([]byte, 8)
	rand.Read(buf)
	return binary.LittleEndian.Uint64(buf)
}
