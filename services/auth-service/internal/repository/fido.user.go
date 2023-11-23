package repository

import (
	"encoding/binary"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
)

// AuthnUser represents the user model
type AuthnUser struct {
	id          uint64
	name        string
	displayName string
	credentials []webauthn.Credential
}

// WebAuthnID returns the user's ID
func (u *AuthnUser) WebAuthnID() []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	binary.PutUvarint(buf, u.id)
	return buf
}

func (u *AuthnUser) WebAuthnName() string {
	return u.name
}

func (u *AuthnUser) WebAuthnDisplayName() string {
	return u.displayName
}

func (u *AuthnUser) WebAuthnCredentials() []webauthn.Credential {
	return u.credentials
}

// AddCredential associates the credential to the user
func (u *AuthnUser) AddCredential(cred webauthn.Credential) {
	u.credentials = append(u.credentials, cred)
}

func (u *AuthnUser) WebAuthnIcon() string {
	return ""
}

// CredentialExcludeList returns a CredentialDescriptor array filled
// with all a user's credentials
func (u *AuthnUser) CredentialExcludeList() []protocol.CredentialDescriptor {

	var credentialExcludeList []protocol.CredentialDescriptor
	for _, cred := range u.credentials {
		descriptor := protocol.CredentialDescriptor{
			Type:         protocol.PublicKeyCredentialType,
			CredentialID: cred.ID,
		}
		credentialExcludeList = append(credentialExcludeList, descriptor)
	}

	return credentialExcludeList
}

// NewUser creates and returns a new AuthnUser
func NewUser(
	id uint64,
	name string,
	displayName string,
) *AuthnUser {
	user := &AuthnUser{}
	user.id = id
	user.name = name
	user.displayName = displayName
	user.credentials = make([]webauthn.Credential, 0)

	return user
}
