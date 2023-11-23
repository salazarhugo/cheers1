package repository

import (
	"github.com/go-webauthn/webauthn/webauthn"
	"time"
)

// Credential model
type Credential struct {
	ID              []byte `gorm:"primarykey"`
	UserId          int64
	PublicKey       []byte
	AttestationType string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func CredentialToDomain(
	userID int64,
	credential *webauthn.Credential,
) *Credential {
	return &Credential{
		ID:              credential.ID,
		UserId:          userID,
		PublicKey:       credential.PublicKey,
		AttestationType: credential.AttestationType,
	}
}

func (c *Credential) ToWebAuthnCredential() *webauthn.Credential {
	return &webauthn.Credential{
		ID:              c.ID,
		PublicKey:       c.PublicKey,
		AttestationType: c.AttestationType,
		Transport:       nil,
		Flags:           webauthn.CredentialFlags{},
		Authenticator:   webauthn.Authenticator{},
	}
}
